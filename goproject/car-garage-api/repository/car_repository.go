package repository

import (
	"car-garage-api/models"
	"github.com/jinzhu/gorm"
)

type CarRepository struct {
	DB *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{DB: db}
}

func (r *CarRepository) CreateCar(car *models.Car) {
	r.DB.Create(car)
}

func (r *CarRepository) GetCars() []models.Car {
	var cars []models.Car
	r.DB.Find(&cars)
	return cars
}

func (r *CarRepository) UpdateCar(car *models.Car) {
	r.DB.Save(car)
}

func (r *CarRepository) DeleteCar(car *models.Car) {
	r.DB.Delete(car)
}
