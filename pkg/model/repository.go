package model

type Repository struct {
	ID        int64     `gorm:"primary_key"`
	Name      string    `gorm:"unique_index"`
	Branches  []*Branch `gorm:"foreignkey:RepositoryID"`
	ProjectID int64
}
