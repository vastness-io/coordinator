package model

type Branch struct {
	ID           int64 `gorm:"primary_key"`
	Name         string
	Languages    []*Language `gorm:"foreignkey:BranchID"`
	Commits      []*Commit   `gorm:"foreignkey:BranchID"`
	RepositoryID int64
}
