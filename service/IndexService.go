package service

import (
	"gocms/models"
)

func GetCate() ([]models.Category) {
	list := make([]models.Category, 0)
	engine.Find(&list)
	return list
}

func GetGlobalList() []models.Data {
	list := make([]models.Data, 0)
	engine.Limit(100, 0).Find(&list)
	return list
}

func GetList(limits int) []models.Data {
	list := make([]models.Data, 0)
	engine.Where("id>?", 0).Limit(0, limits).Find(&list)
	return list
}
