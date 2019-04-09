package service

import "gocms/models"


func GetCate() ([]models.Category)  {
	list := make([]models.Category, 0)
	engine.Find(&list)
	return list
}


func GetGlobalList() *models.Category {
	list := new(models.Category)
	engine.Where("category_id>?", 0).Get(list)
	return list
}


func GetList(limits int) *models.Category {
	list := new(models.Category)
	engine.Limit(limits).Get(list)
	return list
}
