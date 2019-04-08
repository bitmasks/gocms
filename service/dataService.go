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
