package model

import (
	"github.com/lib/pq"
	"time"
)

type Commit struct {
	Sha         string `gorm:"primary_key;size:255;unique_index"`
	Message     string `gorm:"type:varchar(255)"`
	Timestamp   *time.Time
	AuthorName  string         `gorm:"type:varchar(255)"`
	AuthorEmail string         `gorm:"type:varchar(255)"`
	Added       pq.StringArray
	Modified    pq.StringArray
	Removed     pq.StringArray
	BranchID    int64
}
