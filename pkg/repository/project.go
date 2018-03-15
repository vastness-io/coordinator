package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
)

type projectRepository struct {
	tx *gorm.DB
}

func NewProjectRepository(tx *gorm.DB) ProjectRepository {
	return &projectRepository{
		tx: tx,
	}
}

func (r *projectRepository) DB()  *gorm.DB {
	return r.tx
}

func (r *projectRepository) Create(tx *gorm.DB, project *model.Project) error {
	return tx.Create(project).Error
}

func (r *projectRepository) GetProject(tx *gorm.DB, name string, vcsType string) (*model.Project, error) {
	var out model.Project

	err := tx.First(&out, "name = ? AND type = ?", name, vcsType).Error

	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return &out, errors.ProjectDoesNotExistErr
		}

		return &out, err
	}

	return &out, nil
}

func (r *projectRepository) GetProjects(tx *gorm.DB) ([]*model.Project, error) {
	var out []*model.Project

	err := tx.Find(out).Error

	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return out, nil
		}
		return out, err
	}

	return out, nil
}

func (r *projectRepository) Delete(tx *gorm.DB, name string, vcsType string) (bool, error) {
	toBeDeleted := model.Project{
		Name: name,
		Type: vcsType,
	}
	err := tx.Delete(&toBeDeleted).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *projectRepository) Update(tx *gorm.DB, project *model.Project) error {
	return tx.Save(project).Error
}
