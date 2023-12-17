package api

import (
	"car-garage-api/models"
	"car-garage-api/repository"
	"encoding/json"
	"net/http"
)

type CarHandler struct {
	Repo *repository.CarRepository
}

func NewCarHandler(repo *repository.CarRepository) *CarHandler {
	return &CarHandler{Repo: repo}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Repo.CreateCar(&car)
	json.NewEncoder(w).Encode(car)
}

func (h *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	cars := h.Repo.GetCars()
	json.NewEncoder(w).Encode(cars)
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Repo.UpdateCar(&car)
	json.NewEncoder(w).Encode(car)
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	
	// You can use carID to perform any logic you need, such as converting it to uint, etc.
	// For simplicity, we are using it directly here.

	var car models.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Repo.DeleteCar(&car)
	json.NewEncoder(w).Encode(car)
}
