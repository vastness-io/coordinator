package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
)

type projectRepository struct {
	tx DB
}

func NewProjectRepository(tx DB) ProjectRepository {
	return &projectRepository{
		tx: tx,
	}
}

func (r *projectRepository) DB() DB {
	return r.tx
}

func (r *projectRepository) Create(tx DB, project *model.Project) error {
	return tx.Create(project)
}

func (r *projectRepository) GetProject(tx DB, name string, vcsType string) (*model.Project, error) {
	var out model.Project

	err := tx.Preload("Repositories.Branches").First(&out, "name = ? AND type = ?", name, vcsType)

	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return &out, errors.ProjectDoesNotExistErr
		}

		return &out, err
	}

	return &out, nil
}

func (r *projectRepository) GetProjects(tx DB) ([]*model.Project, error) {
	var out []*model.Project

	err := tx.Find(&out)

	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return out, nil
		}
		return out, err
	}

	return out, nil
}

func (r *projectRepository) Delete(tx DB, name string, vcsType string) (bool, error) {
	toBeDeleted := model.Project{
		Name: name,
		Type: vcsType,
	}
	err := tx.Delete(&toBeDeleted)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *projectRepository) Update(tx DB, project *model.Project) error {
	return tx.Save(project)
}
