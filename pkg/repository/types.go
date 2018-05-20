package repository

import (
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/gormer"
)

type ProjectRepository interface {
	DB() gormer.DB
	Create(tx gormer.DB, project *model.Project) error
	GetProject(tx gormer.DB, name string, vcsType string) (*model.Project, error)
	GetProjects(tx gormer.DB, offset, limit int) (*model.ProjectPage, error)
	Delete(tx gormer.DB, name string, vcsType string) (bool, error)
	Update(tx gormer.DB, project *model.Project) error
}
