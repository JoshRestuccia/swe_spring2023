package investments

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3306)/portfullio?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InitialMigration() {
	DB, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

func GetUsersStocks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	wd := uint(id)
	var stocks []Stock
	var names []string
	DB.Find(&stocks, "user_refer=?", wd)
	for i := range stocks {
		names = append(names, stocks[i].Name)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&names)
}

func GetUsersTotalcrypto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	wd := uint(id)
	var crypto []crypto
	var total float64
	DB.Find(&crypto, "user_refer=?", wd)
	for i := range crypto {
		total += (float64(crypto[i].Amount) * (crypto[i].DollarConvert))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&total)
}

func GetUsersTotalStocks(w http.ResponseWriter, r *http.Request) {
	//returns total value of all stocks owned by the user
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())
	}

	wd := uint(id)
	var stocks []Stock
	var total float64

	//find all stocks matching the user id
	DB.Find(&stocks, "user_refer=?", wd)
	for i := range stocks {
		total += float64(stocks[i].Quantity) * stocks[i].Price
	}

	fmt.Printf("Total stock portfolio value: $%.2f\n", total)
	json.NewEncoder(w).Encode(total)
}

func GetUsersTotalCash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())
	}

	wd := uint(id)
	var cash []Cash
	var total float64

	DB.Find(&cash, "user_refer=?", wd)
	for i := range cash {
		total += float64(cash[i].Amount)
	}

	fmt.Printf("Total cash assets value: $%.2f\n", total)
	json.NewEncoder(w).Encode(total)
}

func GetUsersTotal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())
	}

	wd := uint(id)
	var cash []Cash
	var stocks []Stock
	var crypto []crypto
	var total, total2, total3 float64

	DB.Find(&cash, "user_refer=?", wd)
	for i := range cash {
		total += float64(cash[i].Amount)
	}

	fmt.Printf("Total cash assets value: $%.2f\n", total)

	DB.Find(&stocks, "user_refer=?", wd)
	for i := range stocks {
		total2 += float64(stocks[i].Quantity) * stocks[i].Price
	}

	fmt.Printf("Total stock portfolio value: $%.2f\n", total2)

	DB.Find(&crypto, "user_refer=?", wd)
	for i := range crypto {
		total3 += float64(crypto[i].Amount) * float64(crypto[i].DollarConvert)
	}

	fmt.Printf("Total crypto portfolio value: $%.2f\n", total3)

	total += total2 + total3
	fmt.Printf("Total portfolio value: $%.2f\n", total)
	json.NewEncoder(w).Encode(total)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	DB.Find(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	//add a new user

	user := new(User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//remove a user

	id := mux.Vars(r)["id"]
	var user User
	if err := DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	if err := DB.Delete(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "User is deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	//update a user

	id := mux.Vars(r)["id"]

	user := new(User)
	if err := DB.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
