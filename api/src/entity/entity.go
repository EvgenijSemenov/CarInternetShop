package entity

type Car struct {
	Id 		int `json:"id"`
	Status 		string `json:"status"`
	Model 		string `json:"model"`
	Age 		int `json:"age"`
	Race 		int `json:"race"`
	Fuel_type 	string `json:"fuel_type"`
	Price 		int64 `json:"price"`
	Description	string `json:"description"`
}

type User struct {
	Id         int `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Pass_hash  string `json:"pass_hash"`
}
