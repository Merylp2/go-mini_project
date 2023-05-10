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

func GetOrder() ([]models.Order, error) {
	var order []models.Order

	if err := config.DB.Find(&order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func GetOrderByUserId(id int) (order []models.Order, err error) {
	err = config.DB.Preload("Movie").Where("user_id = ?", id).Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func GetOrderByUserID(id int) (order models.Order, err error) {
	err = config.DB.Preload("Movie").Where("user_id = ?", id).Find(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func CheckOrder(order models.Order) (models.Order, error) {
	err := config.DB.Where("user_id = ?", order.UserID).First(&order).Error
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
