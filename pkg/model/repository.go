package model

import "github.com/jinzhu/gorm"

type Repository struct {
	Name            string    `gorm:"primary_key"`
	Owner           string    `gorm:"primary_key"`
	Type            string    `gorm:"primary_key"`
	Branches        []*Branch `gorm:"foreignkey:RepositoryName,RepositoryOwner,RepositoryType"`
	RepositoryName  string    `gorm:"-"`
	RepositoryOwner string    `gorm:"-"`
	RepositoryType  string    `gorm:"-"`
	ProjectID       int64
}

func (r *Repository) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("NAME", r.RepositoryName)
	scope.SetColumn("OWNER", r.RepositoryOwner)
	scope.SetColumn("TYPE", r.RepositoryType)
	return nil
}
