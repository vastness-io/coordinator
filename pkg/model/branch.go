package model

type BranchMeta struct {
	ID        int64      `gorm:"primary_key"`
	Languages []Language `gorm:"many2many:branch_languages"`
	BranchID  int64
}

type Branch struct {
	ID           int64 `gorm:"primary_key"`
	Name         string
	Meta         BranchMeta `gorm:"foreignkey:BranchID"`
	Commits      []*Commit  `gorm:"many2many:branch_commits"`
	RepositoryID int64
}
