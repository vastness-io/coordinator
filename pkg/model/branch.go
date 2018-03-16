package model

type BranchMeta struct {
	ID        int64      `gorm:"primary_key"`
	Languages []Language `gorm:"many2many:branch_languages"`
	BranchID  int64
}
type Branch struct {
	ID           int64      `gorm:"primary_key"`
	Meta         BranchMeta `gorm:"foreignkey:BranchID"`
	Name         string
	Commits      []*Commit `gorm:"foreignkey:BranchID"`
	RepositoryID int64
}
