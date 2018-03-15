package model

type Language struct {
	ID         int64  `gorm:"primary_key"`
	Name       string `gorm:"unique_index"`
	Percentage float64
	BranchID   int64
}
