package project

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/coordinator-svc/project"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/coordinator/pkg/repository/mock"
	"github.com/vastness-io/gormer/mock/golang/mock"
	"reflect"
	"testing"
)

func TestGetProject(t *testing.T) {
	tests := []struct {
		inProject           *model.Project
		mockRepositorySetup func(*mock_repository.MockProjectRepository, *mock_gormer.MockDB, *model.Project, error)
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
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_gormer.MockDB, inProject *model.Project, err error) {
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
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_gormer.MockDB, inProject *model.Project, err error) {
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
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_gormer.MockDB, inProject *model.Project, err error) {
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
				mockDb                = mock_gormer.NewMockDB(ctrl)
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
		mockRepositorySetup func(*mock_repository.MockProjectRepository, *mock_gormer.MockDB, *model.ProjectPage, int, int, error)
		err                 error
		in                  *model.ProjectPage
		startPage           int
		limit               int
		expected            *project.GetProjectsResponse
	}{
		{
			in: &model.ProjectPage{
				Meta: struct {
					CurrentPage int
					LastPage    int
					PerPage     int
					TotalCount  int
				}{CurrentPage: 1, LastPage: 1, PerPage: 5, TotalCount: 5},
				Projects: []*model.Project{
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
			},
			err:       errors.ProjectDoesNotExistErr,
			startPage: 0,
			limit:     5,
			expected:  nil,
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_gormer.MockDB, in *model.ProjectPage, startPage, limit int, err error) {
				mockProjectRepository.EXPECT().DB().Return(mockDb)
				mockProjectRepository.EXPECT().GetProjects(mockDb, startPage, limit).Return(nil, err)
			},
		},
		{
			in: &model.ProjectPage{
				Meta: struct {
					CurrentPage int
					LastPage    int
					PerPage     int
					TotalCount  int
				}{CurrentPage: 1, LastPage: 1, PerPage: 5, TotalCount: 5},
				Projects: []*model.Project{
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
			},
			err:       nil,
			startPage: 1,
			limit:     5,
			expected: &project.GetProjectsResponse{
				Meta: &project.GetProjectsResponse_Meta{
					CurrentPage: 1,
					LastPage:    1,
					PerPage:     5,
					TotalCount:  5,
				},
				Projects: []*project.Project{
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
			},
			mockRepositorySetup: func(mockProjectRepository *mock_repository.MockProjectRepository, mockDb *mock_gormer.MockDB, in *model.ProjectPage, startPage, limit int, err error) {
				mockProjectRepository.EXPECT().DB().Return(mockDb)
				mockProjectRepository.EXPECT().GetProjects(mockDb, startPage, limit).Return(in, err)
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
				mockDb                = mock_gormer.NewMockDB(ctrl)
			)
			defer ctrl.Finish()

			test.mockRepositorySetup(mockProjectRepository, mockDb, test.in, test.startPage, test.limit, test.err)

			projects, err := projectSvc.GetProjects(test.startPage, test.limit)

			if err != test.err {
				t.Fatalf("should equal")
			}

			if !reflect.DeepEqual(test.expected, projects) {
				t.Fatalf("Expected %v, got %v", test.expected, projects)
			}
		}()
	}
}
