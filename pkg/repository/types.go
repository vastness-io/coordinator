package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/vastness-io/coordinator/pkg/model"
)

var (
	PrimaryKeyIsSet = errors.New("primary key has been set manually")
)

type ProjectRepository interface {
	DB ()  *gorm.DB
	Create(tx *gorm.DB, project *model.Project) error
	GetProject(tx *gorm.DB, name string, vcsType string) (*model.Project, error)
	GetProjects(tx *gorm.DB) ([]*model.Project, error)
	Delete(tx *gorm.DB, name string, vcsType string) (bool, error)
	Update(tx *gorm.DB, project *model.Project) error
}

/*
type gormDBWrapper struct {
	*gorm.DB
}

func (t *gormDBWrapper) Begin() DB {
	return NewDB(t.DB.Begin())
}

func (t *gormDBWrapper) Rollback() DB {
	t.DB.Rollback()
	return t
}

func (t *gormDBWrapper) Commit() DB {
	t.DB.Commit()
	return t
}

func (t *gormDBWrapper) Find(value interface{}, where ...interface{}) error {
	return t.DB.Find(value, where).Error
}

func (t *gormDBWrapper) First(out interface{}, where ...interface{}) error {
	return t.DB.First(out, where).Error
}

func (t *gormDBWrapper) Create(value interface{}) error {
	return t.DB.Create(value).Error
}

func (t *gormDBWrapper) Save(value interface{}) error {
	return t.DB.Save(value).Error
}

func (t *gormDBWrapper) Delete(value interface{}, where ...interface{}) error {
	return t.DB.Delete(value, where).Error
}

func (t *gormDBWrapper) GetError() error {
	return t.DB.Error
}


func NewDB(db *gorm.DB) DB {
	return &gormDBWrapper{
		DB: db,
	}
}

type DB interface {
	First(interface{},  ...interface{}) error
	Find(interface{},  ...interface{}) error
	Create(interface{}) error
	Save(interface{}) error
	Delete(interface{}, ...interface{}) error
	Rollback() DB
	Begin() DB
	Commit() DB
	GetError() error
}
*/
