package models

import (
	"time"
)

type Category struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	CategoryId   string    `xorm:"comment('分类ID') unique VARCHAR(200)"`
	CategoryName string    `xorm:"comment('分类名称') VARCHAR(500)"`
	CreatedAt    time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"DATETIME"`
}
