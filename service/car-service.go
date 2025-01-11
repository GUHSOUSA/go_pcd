package service

import (
	"github.com/gustasousaghs/carros-pcd-service/config"
	"github.com/gustasousaghs/carros-pcd-service/models"
)

func CreateCar(user *models.Car) error {
	return config.DB.Create(user).Error
}
func GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	err := config.DB.Find(&cars).Error
	return cars, err
}
func UpdateCar(id string, updatedCar models.Car) error {
	var car models.Car
	if err := config.DB.First(&car, id).Error; err != nil {
		return err
	}
	car.Name = updatedCar.Name
	return config.DB.Save(&car).Error
}
func DeleteCar(id string) error {
	return config.DB.Delete(&models.Car{}, id).Error
}
func GetCarById(id string) (models.Car, error) {
	var car models.Car
	err := config.DB.First(&car, id).Error
	return car, err
}
