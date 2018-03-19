package model

import (
	"github.com/lib/pq"
	"time"
)

type Commit struct {
	Sha         string `gorm:"primary_key;size:256"`
	Message     string `gorm:"type:varchar(256)"`
	Timestamp   *time.Time
	AuthorName  string `gorm:"type:varchar(256)"`
	AuthorEmail string `gorm:"type:varchar(256)"`
	Added       pq.StringArray
	Modified    pq.StringArray
	Removed     pq.StringArray
	BranchID    int64
}
