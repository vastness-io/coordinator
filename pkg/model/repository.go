package model

import "github.com/jinzhu/gorm"

type Repository struct {
	ID           int64     `gorm:"primary_key;AUTO_INCREMENT:false"`
	RepositoryID int64     `gorm:"-"`
	Name         string    `gorm:"unique_index"`
	Branches     []*Branch `gorm:"foreignkey:RepositoryID"`
	ProjectID    int64
}

func (r *Repository) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", r.RepositoryID)
	return nil
}
