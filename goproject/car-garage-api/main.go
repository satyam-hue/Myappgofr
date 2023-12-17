package main

import (
	"car-garage-api/api"
	"car-garage-api/repository"
	"car-garage-api/utils"
	"net/http"
)

func main() {
	utils.InitDB()
	defer utils.DB.Close()

	carRepo := repository.NewCarRepository(utils.DB)
	router := api.NewRouter(carRepo)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
