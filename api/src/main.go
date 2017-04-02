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
	"github.com/dgrijalva/jwt-go"
	"crypto/md5"

	"entity"
	"time"
	"errors"
	"response"
)

const (
	DB_USER = "postgres"
	DB_PASSWORD = "secret"
	DB_NAME = "shop_car"

	API_URI = "/api/v1"

)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(API_URI + "/car/all", getAllCars).Methods("GET")
	router.HandleFunc(API_URI + "/car/{carId}", getCar).Methods("GET")
	router.HandleFunc(API_URI + "/car/add", addCar).Methods("POST")
	router.HandleFunc(API_URI + "/car/update", updateCar).Methods("POST")
	router.HandleFunc(API_URI + "/car/delete", deleteCar).Methods("POST")

	router.HandleFunc(API_URI + "/user/add", addUser).Methods("POST")
	router.HandleFunc(API_URI + "/user/login", loginUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDBConn() sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	return *db;
}

func getAllCars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Get all cars")

	var db sql.DB = getDBConn()

	rows, err := db.Query("SELECT * FROM car")
	checkErr(err)

	var cars []entity.Car

	for rows.Next() {
		var car entity.Car
		var description sql.NullString
		var price sql.NullInt64

		err = rows.Scan(&car.Id, &car.Status, &car.Model, &car.Age, &car.Race, &car.Fuel_type, &description, &price)
		checkErr(err)

		car.Price = price.Int64
		car.Description = description.String
		cars = append(cars, car)
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(cars)
}

func getCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Get car")

	var db sql.DB = getDBConn()

	vars := mux.Vars(r)
	getCarId := vars["carId"]

	rows, err := db.Query("SELECT * FROM car WHERE id=" + getCarId)
	checkErr(err)

	var cars []entity.Car

	for rows.Next() {
		var car entity.Car
		var description sql.NullString
		var price sql.NullInt64

		err = rows.Scan(&car.Id, &car.Status, &car.Model, &car.Age, &car.Race, &car.Fuel_type, &description, &price)
		checkErr(err)

		car.Price = price.Int64
		car.Description = description.String
		cars = append(cars, car)
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(cars)
}

func addCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Add car")

	var db sql.DB = getDBConn()
	var car entity.Car = getCarFromRequest(r)

	var lastInsertId int

	err := db.QueryRow("INSERT INTO car(status, model, age, race, fuel_type, price, description) VALUES ($1,$2,$3,$4,$5,$6,$7) returning id;",
		car.Status, car.Model, car.Age, car.Race, car.Fuel_type, car.Price, car.Description).Scan(&lastInsertId)
	checkErr(err)

	car.Id = lastInsertId

	w.Header().Add("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(car)
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Delete car")

	var db sql.DB = getDBConn()

	stmt, err := db.Prepare("delete from car where id=$1")
	checkErr(err)

	car := getCarFromRequest(r)

	res, err := stmt.Exec(car.Id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE")

	json.NewEncoder(w).Encode("OK")
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Update car")

	var db sql.DB = getDBConn()

	stmt, err := db.Prepare("UPDATE car SET status=$2, model=$3, age=$4, race=$5, fuel_type=$6, price=$7, description=$8 WHERE id=$1")
	checkErr(err)

	car := getCarFromRequest(r)

	fmt.Println(car.Id)

	res, err := stmt.Exec(car.Id, car.Status, car.Model, car.Age, car.Race, car.Fuel_type, car.Price, car.Description)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect, "rows changed")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PUT")

	json.NewEncoder(w).Encode("OK")
}

func getCarFromRequest(r *http.Request) entity.Car {
	var car entity.Car

	var parsed map[string]interface{}

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("getentity.CarFromRequest body:" + string(body))
	checkErr(err)

	err = json.Unmarshal(body, &parsed)
	checkErr(err)

	if (parsed["id"] != nil) {
		car.Id = int(parsed["id"].(float64))
	}
	if (parsed["status"] != nil) {
		car.Status = parsed["status"].(string)
	}
	if (parsed["model"] != nil) {
		car.Model = parsed["model"].(string)
	}
	if (parsed["fuel_type"] != nil) {
		car.Fuel_type = parsed["fuel_type"].(string)
	}
	if (parsed["description"] != nil) {
		car.Description = parsed["description"].(string)
	}
	if (parsed["age"] != nil) {
		car.Age = int(parsed["age"].(float64))
	}
	if (parsed["race"] != nil) {
		car.Race = int(parsed["race"].(float64))
	}
	if (parsed["price"] != nil) {
		car.Price = int64(parsed["price"].(float64))
	}

	return car
}

func addUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Add user")

	var db sql.DB = getDBConn()
	var user entity.User = getUserFromRequest(r)

	var lastInsertId int

	err := db.QueryRow("INSERT INTO \"user\"(First_name, Last_name, Email, Pass_hash) VALUES ($1, $2, $3, $4) returning id;",
		user.First_name, user.Last_name, user.Email, user.Pass_hash).Scan(&lastInsertId)
	checkErr(err)

	user.Id = lastInsertId

	w.Header().Add("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(user)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("# Login user")

	var db sql.DB = getDBConn()
	var user entity.User = getUserFromRequest(r)

	rows, err := db.Query("SELECT first_name, last_name, email, pass_hash FROM \"user\" WHERE email='" + user.Email + "'" )

	var tokenString string
	if rows.Next() {
		var firstName string
		var lastName string
		var email string
		var passHash string

		err = rows.Scan(&firstName, &lastName, &email, &passHash)
		checkErr(err)

		hash := md5.New()
		b := []byte(user.Email + ":" + user.Pass)
		hash.Write(b)
		str := fmt.Sprintf("%x", hash.Sum(nil))
		fmt.Println(str)

		if str == passHash {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"first_name": firstName,
				"last_name": lastName,
				"email": email,
				"exp": time.Now().Add(time.Hour * 12).Unix(),
			})

			tokenString, err = token.SignedString([]byte("secret"))
			fmt.Println("secc")
		} else {
			fmt.Println("err")
			err = errors.New("Incorrect email or password")
		}
	} else {
		err = errors.New("User not found")
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")

	result := make(map[string]interface{})
	result["auth-token"] = tokenString

	json.NewEncoder(w).Encode(response.BuildResponse(result, err))
}

func getUserFromRequest(r *http.Request) entity.User {
	var user entity.User

	var parsed map[string]interface{}

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("getUserFromRequest body:" + string(body))
	checkErr(err)

	err = json.Unmarshal(body, &parsed)
	checkErr(err)

	if (parsed["id"] != nil) {
		user.Id = int(parsed["id"].(float64))
	}
	if (parsed["first_name"] != nil) {
		user.First_name = parsed["first_name"].(string)
	}
	if (parsed["last_name"] != nil) {
		user.Last_name = parsed["last_name"].(string)
	}
	if (parsed["email"] != nil) {
		user.Email = parsed["email"].(string)
	}
	if (parsed["pass_hash"] != nil) {
		user.Pass_hash = parsed["pass_hash"].(string)
	}
	if (parsed["pass"] != nil) {
		user.Pass = parsed["pass"].(string)
	}

	return user
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
