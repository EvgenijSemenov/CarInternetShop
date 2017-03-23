package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"

type Car struct {
	Id int `json:"id"`
	Status string `json:"status"`
	Model string `json:"model"`
	Age int `json:"age"`
	Race int `json:"race"`
	Fuel_type string `json:"fuel_type"`
	Description string `json:"description"`
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "secret"
	DB_NAME     = "shop_car"

	API_URI = "/api/v1/"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(API_URI + "car", getCars)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getCars(w http.ResponseWriter, r *http.Request) {
