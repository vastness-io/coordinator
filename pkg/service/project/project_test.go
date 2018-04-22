package project

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/coordinator-svc/project"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/coordinator/pkg/repository/mock"
	"reflect"
	"testing"
)

func TestGetProject(t *testing.T) {
	tests := []struct {
		inProject           *model.Project
		mockRepositorySetup func(*mock_repository.MockProjectRepository, *mock_repository.MockDB, *model.Project, error)
		err                 error
		expected            *project.Project
	}{
		{
			inProject: &model.Project{
				Name: "project_1",
				Type: "GITHUB",
			},
			err: nil,
			expected: &project.Project{
				Name:         "project_1",
				Type:         "GITHUB",
				Repositories: []*project.Repository{},
			},
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, inProject *model.Project, err error) {
				mockProjectRepository.EXPECT().DB().Return(mockDb)
				mockProjectRepository.EXPECT().GetProject(gomock.Eq(mockDb), inProject.Name, inProject.Type).Return(inProject, err)

			},
		},
		{
			inProject: &model.Project{
				Name: "project_doesn't_exist",
				Type: "GITHUB",
			},
			err:      errors.ProjectDoesNotExistErr,
			expected: nil,
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, inProject *model.Project, err error) {
				mockProjectRepository.EXPECT().DB().Return(mockDb)
				mockProjectRepository.EXPECT().GetProject(gomock.Eq(mockDb), inProject.Name, inProject.Type).Return(nil, err)

			},
		},
		{
			inProject: &model.Project{
				Name: "project_2",
				Type: "BITBUCKET-SERVER",
			},
			err: nil,
			expected: &project.Project{
				Name:         "project_2",
				Type:         "BITBUCKET-SERVER",
				Repositories: []*project.Repository{},
			},
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, inProject *model.Project, err error) {
				mockProjectRepository.EXPECT().DB().Return(mockDb)
				mockProjectRepository.EXPECT().GetProject(gomock.Eq(mockDb), inProject.Name, inProject.Type).Return(inProject, err)

			},
		},
	}

	for _, test := range tests {
		func() {
			var (
				ctrl                  = gomock.NewController(t)
				mockProjectRepository = mock_repository.NewMockProjectRepository(ctrl)
				log                   = logrus.New().WithField("testing", true)
				projectSvc            = NewProjectService(log, mockProjectRepository)
				mockDb                = mock_repository.NewMockDB(ctrl)
			)
			defer ctrl.Finish()

			test.mockRepositorySetup(mockProjectRepository, mockDb, test.inProject, test.err)

			p, err := projectSvc.GetProject(test.inProject.Name, test.inProject.Type)

			if err != test.err {
				t.Fatalf("Expected %v, got %v", test.err, err)
			}

			if !reflect.DeepEqual(test.expected, p) {
				t.Fatalf("Expected %v, got %v", p, test.expected)
			}
		}()
	}
}

func TestGetProjects(t *testing.T) {
	tests := []struct {
		mockRepositorySetup func(*mock_repository.MockProjectRepository, *mock_repository.MockDB, []*model.Project, error)
		err                 error
		in                  []*model.Project
		expected            []*project.Project
		error               error
	}{
		{
			in: []*model.Project{
				{
					Name:         "project_1",
					Type:         "GITHUB",
					Repositories: []*model.Repository{},
				},
				{
					Name:         "project_2",
					Type:         "GITHUB",
					Repositories: []*model.Repository{},
				},
				{
					Name:         "project_3",
					Type:         "BITBUCKET_SERVER",
					Repositories: []*model.Repository{},
				},
			},
			err:      errors.ProjectDoesNotExistErr,
			expected: nil,
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, in []*model.Project, err error) {
				mockProjectRepository.EXPECT().DB().Return(mockDb)
				mockProjectRepository.EXPECT().GetProjects(gomock.Eq(mockDb)).Return(nil, err)

			},
		},
		{
			in: []*model.Project{
				{
					Name:         "project_1",
					Type:         "GITHUB",
					Repositories: []*model.Repository{},
				},
				{
					Name:         "project_2",
					Type:         "GITHUB",
					Repositories: []*model.Repository{},
				},
				{
					Name:         "project_3",
					Type:         "BITBUCKET_SERVER",
					Repositories: []*model.Repository{},
				},
			},
			err: nil,
			expected: []*project.Project{
				{
					Name:         "project_1",
					Type:         "GITHUB",
					Repositories: []*project.Repository{},
				},
				{
					Name:         "project_2",
					Type:         "GITHUB",
					Repositories: []*project.Repository{},
				},
				{
					Name:         "project_3",
					Type:         "BITBUCKET_SERVER",
					Repositories: []*project.Repository{},
				},
			},
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_repository.MockDB, in []*model.Project, err error) {
				mockProjectRepository.EXPECT().DB().Return(mockDb)
				mockProjectRepository.EXPECT().GetProjects(gomock.Eq(mockDb)).Return(in, err)

			},
		},
	}

	for _, test := range tests {
		func() {
			var (
				ctrl                  = gomock.NewController(t)
				mockProjectRepository = mock_repository.NewMockProjectRepository(ctrl)
				log                   = logrus.New().WithField("testing", true)
				projectSvc            = NewProjectService(log, mockProjectRepository)
				mockDb                = mock_repository.NewMockDB(ctrl)
			)
			defer ctrl.Finish()

			test.mockRepositorySetup(mockProjectRepository, mockDb, test.in, test.error)

			projects := projectSvc.GetProjects()

			if !reflect.DeepEqual(test.expected, projects) {
				t.Fatalf("Expected %v, got %v", test.expected, projects)
			}
		}()
	}
}
