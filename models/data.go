package models

import (
	"time"
)

type Data struct {
	Id           int64     `xorm:"pk autoincr BIGINT(20)"`
	Spider       string    `xorm:"comment('爬虫：名称') VARCHAR(200)"`
	OriginUrl    string    `xorm:"comment('爬虫：来源url') VARCHAR(1000)"`
	CategoryName string    `xorm:"default '' comment('文章分类名称') VARCHAR(255)"`
	CategoryId   string    `xorm:"comment('文章分类ID') VARCHAR(200)"`
	Title        string    `xorm:"comment('文章分类名称') VARCHAR(500)"`
	Cover        string    `xorm:"comment('文章封面图') VARCHAR(1000)"`
	Content      string    `xorm:"comment('文章内容') LONGTEXT"`
	Author       string    `xorm:"comment('文章作者') VARCHAR(255)"`
	Desc         string    `xorm:"default '' comment('文章简介') VARCHAR(500)"`
	Tag          string    `xorm:"default '' comment('文章关键词') VARCHAR(500)"`
	CreatedAt    time.Time `xorm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"TIMESTAMP"`
}
