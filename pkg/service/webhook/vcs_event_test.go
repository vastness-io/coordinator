package webhook

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	errors2 "github.com/vastness-io/coordinator/pkg/errors"
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
			err: errors.New("example"),
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
		expectedDetectedLangs []*linguist.Language
		err                   error
		project               *model.Project
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
			project: &model.Project{
				Name: "example",
				Type: "GITHUB",
			},
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

			mockProjectRepository.EXPECT().DB().Return(mockDb)

			mockDb.EXPECT().Begin().Return(mockDb)

			mockDb.EXPECT().GetError().Return(nil)

			mockProjectRepository.EXPECT().GetProject(gomock.Eq(mockDb), test.project.Name, test.project.Type).Return(nil, errors2.ProjectDoesNotExistErr)

			mockProjectRepository.EXPECT().Create(gomock.Eq(mockDb), test.project).Return(nil)

			mockProjectRepository.EXPECT().Update(gomock.Eq(mockDb), test.project).Return(nil)

			mockDb.EXPECT().Commit().Return(mockDb)

			mockDb.EXPECT().GetError().Return(nil)

			p, err := eventSvc.UpdateProject(test.project)

			if err != nil {
				t.Fatal("Update should have succeeded")
			}

			if !reflect.DeepEqual(test.project, p) {
				t.Fatalf("Expected %v, got %v", p, test.project)
			}

		}()
	}
}
