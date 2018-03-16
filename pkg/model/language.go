package model

type Language struct {
	ID         int64 `gorm:"primary_key"`
	Name       string
	Percentage float64
}
