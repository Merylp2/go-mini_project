package database

import (
	"go_movie-ticket/config"
	"go_movie-ticket/models"
)

func CreateOrder(order models.Order) (models.Order, error) {
	err := config.DB.Create(&order).Error

	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func GetOrder() (order models.Order, err error) {
	err = config.DB.First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
