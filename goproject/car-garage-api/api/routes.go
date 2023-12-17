package api

import (
    "car-garage-api/repository"
    "github.com/gorilla/mux"
)

func NewRouter(repo *repository.CarRepository) *mux.Router {
    r := mux.NewRouter()
    carHandler := NewCarHandler(repo)

    r.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
    r.HandleFunc("/cars", carHandler.GetCars).Methods("GET")
    r.HandleFunc("/cars", carHandler.UpdateCar).Methods("PUT")
    r.HandleFunc("/cars", carHandler.DeleteCar).Methods("DELETE")

    return r
}
