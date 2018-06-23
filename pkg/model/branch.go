package model

type Branch struct {
	Meta            *BranchMeta
	ID              int64 `gorm:"primary_key"`
	Name            string
	Commits         []*Commit `gorm:"many2many:branch_commits"`
	RepositoryName  string
	RepositoryOwner string
	RepositoryType  string
}
