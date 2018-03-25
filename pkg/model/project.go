package model

type Project struct {
	ID           int64         `gorm:"primary_key"`
	Name         string        `gorm:"unique_index"`
	Type         string        `gorm:"unique_index"`
	Repositories []*Repository `gorm:"foreignkey:ProjectID"`
}
