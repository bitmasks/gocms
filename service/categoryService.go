package service

import "gocms/models"

func GetCategorys() ([]models.Category)  {
	list := make([]models.Category, 0)
	engine.Find(&list)
	return list
}

func GetCategoryInfo(ctId string) *models.Category {
	info := new(models.Category)
	engine.Where("category_id=?", ctId).Get(info)
	return info
}
