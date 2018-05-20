package project

import (
	"github.com/sirupsen/logrus"
	"github.com/vastness-io/coordinator-svc/project"
	"github.com/vastness-io/coordinator/pkg/repository"
)

type projectService struct {
	logger     *logrus.Entry
	repository repository.ProjectRepository
}

func NewProjectService(logger *logrus.Entry, projectRepository repository.ProjectRepository) Service {
	return &projectService{
		logger:     logger,
		repository: projectRepository,
	}
}

func (ps *projectService) GetProject(name, vcsType string) (*project.Project, error) {
	db := ps.repository.DB()
	projectModel, err := ps.repository.GetProject(db, name, vcsType)

	if err != nil {
		return nil, err
	}

	return FromProjectModel(projectModel), nil

}

func (ps *projectService) GetProjects(offset, limit int) (*project.GetProjectsResponse, error) {
	db := ps.repository.DB()
	projectPage, err := ps.repository.GetProjects(db, offset, limit)

	if err != nil {
		return nil, err
	}

	out := project.GetProjectsResponse{
		Projects: make([]*project.Project, len(projectPage.Projects)),
	}

	for i, _ := range projectPage.Projects {
		out.Projects[i] = FromProjectModel(projectPage.Projects[i])
	}

	out.Meta = &project.GetProjectsResponse_Meta{
		CurrentPage: int32(projectPage.Meta.CurrentPage),
		LastPage:    int32(projectPage.Meta.LastPage),
		PerPage:     int32(projectPage.Meta.PerPage),
		TotalCount:  int32(projectPage.Meta.TotalCount),
	}

	return &out, nil
}
