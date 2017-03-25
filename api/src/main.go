package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
)

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

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM car")
	checkErr(err)

	var cars []Car

	for rows.Next() {
		var car Car
		var description sql.NullString

		err = rows.Scan(&car.Id, &car.Status, &car.Model, &car.Age, &car.Race, &car.Fuel_type, &description)
		checkErr(err)

		car.Description = description.String
		cars = append(cars, car)
	}

	json.NewEncoder(w).Encode(cars)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
