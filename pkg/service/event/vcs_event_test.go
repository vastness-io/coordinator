package event

import (
	stdlib "errors"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/coordinator/pkg/repository/mock"
	"github.com/vastness-io/linguist-svc"
	"github.com/vastness-io/linguist-svc/mock/linguist"
	"reflect"
	"testing"
)

func TestGetLanguagesUsedInRepository(t *testing.T) {

	tests := []struct {
		expectedDetectedLangs []*linguist.Language
		err                   error
	}{
		{
			expectedDetectedLangs: []*linguist.Language{
				{
					Name:       "Go",
					Percentage: float64(50),
				},
				{
					Name:       "Java",
					Percentage: float64(50),
				},
			},
		},
		{

			expectedDetectedLangs: []*linguist.Language{},
			err: stdlib.New("example"),
		},
	}

	for _, test := range tests {

		func() {
			var (
				ctrl                  = gomock.NewController(t)
				mockClient            = mock_linguist_svc.NewMockLinguistClient(ctrl)
				mockProjectRepository = mock_repository.NewMockProjectRepository(ctrl)
				log                   = logrus.New().WithField("testing", true)
				eventSvc              = NewVcsEventService(log, mockClient, mockProjectRepository)
				langReq               = &linguist.LanguageRequest{
					FileNames: []string{
						"main.go",
						"app.java",
					},
				}

				langRes = &linguist.LanguageResponse{
					test.expectedDetectedLangs,
				}
			)
			defer ctrl.Finish()

			mockClient.EXPECT().GetLanguages(gomock.Any(), langReq).Return(langRes, test.err)

			languages := eventSvc.GetLanguagesUsedInRepository(langReq)

			if len(languages) != len(test.expectedDetectedLangs) {
				t.Fatal("Should equal")
			}

			for _, expected := range test.expectedDetectedLangs {
				var found bool
				for _, actual := range languages {
					if expected == actual {
						found = true
					}
				}
				if !found {
					t.Fatalf("Should have detected %s language", expected)
				}

				found = false
			}
		}()
	}
}

func TestUpdateProject(t *testing.T) {

	tests := []struct {
		inProject  *model.Project
		outProject *model.Project
		err        error
		mockSetup  func(*mock_repository.MockProjectRepository, *mock_repository.MockDB, *mock_linguist_svc.MockLinguistClient, *model.Project, error)
	}{
		{
			inProject: &model.Project{
				Name: "example",
				Type: "GITHUB",
			},
			outProject: nil,

			err: stdlib.New("testing handle of unexpected error returned from db"),

			mockSetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, mockLing *mock_linguist_svc.MockLinguistClient, project *model.Project, err error) {

				mockProjectRepository.EXPECT().DB().Return(mockDb)

				mockDb.EXPECT().Begin().Return(mockDb)

				mockDb.EXPECT().GetError().Return(nil)

				mockProjectRepository.EXPECT().GetProject(gomock.Eq(mockDb), project.Name, project.Type).Return(nil, err)
			},
		},
		{
			mockSetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, mockLing *mock_linguist_svc.MockLinguistClient, project *model.Project, err error) {
				// This test checks if project exists, and performs creation if necessary
				mockProjectRepository.EXPECT().DB().Return(mockDb)

				mockDb.EXPECT().Begin().Return(mockDb)

				mockDb.EXPECT().GetError().Return(nil)

				mockProjectRepository.EXPECT().GetProject(gomock.Eq(mockDb), project.Name, project.Type).Return(nil, errors.ProjectDoesNotExistErr)

				mockProjectRepository.EXPECT().Create(gomock.Eq(mockDb), project).Return(nil)

				mockProjectRepository.EXPECT().Update(gomock.Eq(mockDb), project).Return(nil)

				req := linguist.LanguageRequest{
					FileNames: []string{
						"main.go",
						"build.xml",
					},
				}

				mockLing.EXPECT().GetLanguages(gomock.Any(), &req)

				mockDb.EXPECT().Commit().Return(mockDb)

				mockDb.EXPECT().GetError().Return(err)

			},
			err: nil,
			inProject: &model.Project{
				Name: "example",
				Type: "GITHUB",
				Repositories: []*model.Repository{
					{
						Name: "repo_1",
						Branches: []*model.Branch{
							{
								Name: "branch_1",
								Commits: []*model.Commit{
									{
										Added: []string{
											"main.go",
										},
										Modified: []string{
											"build.xml",
										},
									},
								},
							},
						},
					},
				},
			},
			outProject: &model.Project{
				Name: "example",
				Type: "GITHUB",
				Repositories: []*model.Repository{
					{
						Name: "repo_1",
						Branches: []*model.Branch{
							{
								Name: "branch_1",
								Commits: []*model.Commit{
									{
										Added: []string{
											"main.go",
										},
										Modified: []string{
											"build.xml",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			mockSetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, mockLing *mock_linguist_svc.MockLinguistClient, project *model.Project, err error) {
				// Test for rollback
				mockProjectRepository.EXPECT().DB().Return(mockDb)

				mockDb.EXPECT().Begin().Return(mockDb)

				mockDb.EXPECT().GetError().Return(nil)

				mockProjectRepository.EXPECT().GetProject(gomock.Eq(mockDb), project.Name, project.Type).Return(nil, errors.ProjectDoesNotExistErr)

				mockProjectRepository.EXPECT().Create(gomock.Eq(mockDb), project).Return(err)

				mockDb.EXPECT().Rollback()

			},
			err: stdlib.New("unexpected error to cause rollback"),

			inProject: &model.Project{
				Name: "example",
				Type: "GITHUB",
			},
			outProject: nil,
		},
	}

	for _, test := range tests {

		func() {
			var (
				ctrl                  = gomock.NewController(t)
				mockClient            = mock_linguist_svc.NewMockLinguistClient(ctrl)
				mockProjectRepository = mock_repository.NewMockProjectRepository(ctrl)
				log                   = logrus.New().WithField("testing", true)
				eventSvc              = NewVcsEventService(log, mockClient, mockProjectRepository)
				mockDb                = mock_repository.NewMockDB(ctrl)
			)
			defer ctrl.Finish()

			test.mockSetup(mockProjectRepository, mockDb, mockClient, test.inProject, test.err)

			p, err := eventSvc.UpdateProject(test.inProject)

			if err != test.err {
				t.Fatalf("Expected %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(test.outProject, p) {
				t.Fatalf("Expected %v, got %v", p, test.outProject)
			}

		}()
	}
}
