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

)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}