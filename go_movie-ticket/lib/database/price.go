package database

import (
	"go_movie-ticket/config"
	"go_movie-ticket/models"
)

func CreatePrice(price models.Price) (models.Price, error) {
	err := config.DB.Create(&price).Error

	if err != nil {
		return models.Price{}, err
	}
	return price, nil
}

func UpdatePrice(price models.Price, id any) (models.Price, error) {
	err := config.DB.Table("prices").Where("id = ?", id).Updates(&price).Error

	if err != nil {
		return models.Price{}, err
	}

	return price, nil
}

func DeletePrice(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.Price{}).Error

	if err != nil {
		return nil, err
	}

	return "success delete price", nil
}