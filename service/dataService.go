package service

import "gocms/models"

func GetDataInfo(id int64) (*models.Data ,error)   {
	data := new(models.Data)
	has, err := engine.Id(id).Get(data)
	if has {
		return data , nil
	}
	return nil ,err
}

func GetDataList(cId string) []models.Data {
	list := make([]models.Data, 0)
	engine.Where("category_id=?",cId).Find(&list)
	return list
}