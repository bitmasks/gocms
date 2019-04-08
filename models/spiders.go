package models

import (
	"time"
)

type Spiders struct {
	Id                int       `xorm:"not null pk autoincr INT(11)"`
	Name              string    `xorm:"default '' unique VARCHAR(255)"`
	Status            int       `xorm:"default 0 TINYINT(1)"`
	ScanUrls          string    `xorm:"VARCHAR(255)"`
	ContentUrlRegexes string    `xorm:"VARCHAR(255)"`
	ListUrlRegexes    string    `xorm:"VARCHAR(5000)"`
	CategoryId        string    `xorm:"VARCHAR(200)"`
	CategoryName      string    `xorm:"default '' VARCHAR(200)"`
	Fields            string    `xorm:"VARCHAR(5000)"`
	CreatedAt         time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt         time.Time `xorm:"TIMESTAMP"`
}
