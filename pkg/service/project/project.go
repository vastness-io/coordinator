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

func (ps *projectService) GetProjects() (out []*project.Project) {
	db := ps.repository.DB()
	projectModels, err := ps.repository.GetProjects(db)

	if err != nil {
		return
	}

	for _, prj := range projectModels {
		out = append(out, FromProjectModel(prj))
	}
	return
}
