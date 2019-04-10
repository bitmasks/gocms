package service

import (
	"gocms/models"
)

func GetCate() ([]models.Category) {
	list := make([]models.Category, 0)
	engine.Find(&list)
	return list
}

func GetGlobalList(limits int) []models.Data {
	list := make([]models.Data, 0)
	engine.Limit(limits, 0).Find(&list)
	return list
}

func GetList( cid  string , limits int) []models.Data {
	list := make([]models.Data, 0)
	engine.Where("category_id =?", cid).Limit(limits,0).Find(&list)
	return list
}
