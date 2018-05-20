package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/vastness-io/coordinator/pkg/errors"
	"github.com/vastness-io/coordinator/pkg/model"
	"github.com/vastness-io/gormer"
)

type projectRepository struct {
	tx gormer.DB
}

func NewProjectRepository(tx gormer.DB) ProjectRepository {
	return &projectRepository{
		tx: tx,
	}
}

func (r *projectRepository) DB() gormer.DB {
	return r.tx
}

func (r *projectRepository) Create(tx gormer.DB, project *model.Project) error {
	return tx.Create(project).Error()
}

func (r *projectRepository) GetProject(tx gormer.DB, name string, vcsType string) (*model.Project, error) {
	var out model.Project

	err := tx.Preload("Repositories.Branches").Preload("Repositories.Branches.Commits").First(&out, "name = ? AND type = ?", name, vcsType).Error()

	if err != nil {

		if gorm.IsRecordNotFoundError(err) {
			return &out, errors.ProjectDoesNotExistErr
		}

		return &out, err
	}

	return &out, nil
}

func (r *projectRepository) GetProjects(tx gormer.DB, offset, limit int) (*model.ProjectPage, error) {

	var projects []*model.Project

	santisedOffset := offset

	if offset == 0 {
		santisedOffset = 1
	}

	stmt := tx.
		Preload("Repositories.Branches").
		Preload("Repositories.Branches.Commits")

	if limit != 0 {
		stmt = stmt.Limit(limit)
	}

	var count int

	err := stmt.Find(&projects).Count(&count).Error()

	if err != nil {
		return nil, err
	}

	out := model.ProjectPage{
		Meta: struct {
			CurrentPage int
			LastPage    int
			PerPage     int
			TotalCount  int
		}{
			CurrentPage: santisedOffset,
			LastPage:    GetTotalPages(count, limit),
			PerPage:     limit,
			TotalCount:  GetTotalPages(count, limit),
		},
		Projects: projects,
	}

	return &out, nil
}

func (r *projectRepository) Delete(tx gormer.DB, name string, vcsType string) (bool, error) {
	toBeDeleted := model.Project{
		Name: name,
		Type: vcsType,
	}
	err := tx.Delete(&toBeDeleted).Error()

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *projectRepository) Update(tx gormer.DB, project *model.Project) error {
	return tx.Save(project).Error()
}

func GetTotalPages(count int, perPage int) int {
	totalPages := float32(count) / float32(perPage)

	return int(totalPages + 1.0)
}
