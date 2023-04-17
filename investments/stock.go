package investments

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Stock struct {
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity" gorm:"default:1"`
	Favorite  bool    `json:"favorite" gorm:"default:false"`
	UserRefer uint    `json:"userRefer"`
	//	User      User    `gorm:"foreignKey:UserRefer"`
}

func MigrateStocks() {
	//initialize stock table

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	DB.AutoMigrate(&Stock{})
}

func FindUser(id uint, user *User) User {
	//finds the user who has the stock we are trying to add

	DB.Find(&user, "id =?", id)

	if user.ID == id {
		return *user
	}
	return User{}
}

func CreateResponseStock(stock Stock) Stock {
	return Stock{Symbol: stock.Symbol, Name: stock.Name, Price: stock.Price, Quantity: stock.Quantity, UserRefer: stock.UserRefer}
}

func GetStocks(w http.ResponseWriter, r *http.Request) {
	//returns all stocks of a given user
	params := mux.Vars(r)
	user_refer := params["user_refer"]

	//get the user id number as a string
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stocks []Stock
	//find all stocks matching the user id
	DB.Find(&stocks, "user_refer=?", wd)

	json.NewEncoder(w).Encode(stocks)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_refer := params["user_refer"]
	symbol := params["symbol"]
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert id to uint

	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stock Stock
	DB.Where("user_refer=?", wd).Where("symbol=?", symbol).Find(&stock)
	json.NewEncoder(w).Encode(stock)
}

func SaveStock(w http.ResponseWriter, r *http.Request) {
	//Should be added to add stock with user_refer
	//adds a new stock
	stock := new(Stock)
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//var user User

	DB.Create(&stock)
	//stock.User = FindUser(stock.UserRefer, &user)

	json.NewEncoder(w).Encode(stock)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	//removes a single stock from a user's portfolio
	vars := mux.Vars(r)
	userRefer := vars["user_refer"]
	symbol := vars["symbol"]

	// get the user id number and stock symbol as strings
	u64, err := strconv.ParseUint(userRefer, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	wd := uint(u64)

	var stock Stock
	//delete stocks matching the user id and stock symbol
	DB.Where("user_refer=?", wd).Where("symbol=?", symbol).Unscoped().Delete(&stock)

	fmt.Fprint(w, "Stock deleted")
}

func DeleteStocks(w http.ResponseWriter, r *http.Request) {
	//same logic as DeleteStock() but no symbol parameter
	vars := mux.Vars(r)
	userRefer := vars["user_refer"]

	u64, err := strconv.ParseUint(userRefer, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	wd := uint(u64)

	var stock Stock
	DB.Where("user_refer=?", wd).Delete(&stock)

	fmt.Fprint(w, "Stock deleted")
}

func findStock(symbol string, id uint, stock Stock) error {
	DB.Where("symbol=?", symbol).Where("user_refer=?", id).Find(&stock)
	if stock.Symbol == "" {
		return errors.New("Stock not found")
	}
	return nil
}

func ReturnStock(symbol string, id uint, stock Stock) Stock {
	DB.Where("symbol=?", symbol).Where("user_refer=?", id).Find(&stock)
	return stock
}

func GetFavs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userRefer := vars["user_refer"]
	u64, err := strconv.ParseUint(userRefer, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	wd := uint(u64)
	var stocks []Stock
	DB.Where("user_refer=?", wd).Find(&stocks)
	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	//updates a stock
	params := mux.Vars(r)
	symbol := params["symbol"]
	user_refer := params["user_refer"]
	var stock Stock
	u64, er := strconv.ParseUint(user_refer, 10, 32)
	//convert to uint
	if er != nil {
		fmt.Println(er.Error())
	}
	wd := uint(u64)

	err := findStock(symbol, wd, stock)
	stock = ReturnStock(symbol, wd, stock)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Stock not found"))
		return
	}

	type UpdateStock struct {
		Symbol   string  `json:"symbol"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity" gorm:"default:1"`
	}

	updatedInfo := new(UpdateStock)
	err = json.NewDecoder(r.Body).Decode(updatedInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't create new stock"))
		return
	}
	stock.Symbol = updatedInfo.Symbol
	stock.Name = updatedInfo.Name
	stock.Quantity = updatedInfo.Quantity
	stock.Price = updatedInfo.Price

	DB.Where("symbol=?", symbol).Where("user_refer=?", user_refer).Save(&stock)

	json.NewEncoder(w).Encode(stock)
}

func FavoriteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	symbol := params["symbol"]
	user_refer := params["user_refer"]
	var stock Stock
	u64, er := strconv.ParseUint(user_refer, 10, 32)
	//convert to uint
	if er != nil {
		fmt.Println(er.Error())
	}
	wd := uint(u64)

	err := findStock(symbol, wd, stock)
	stock = ReturnStock(symbol, wd, stock)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Stock not found"))
		return
	}

	type UpdateStock struct {
		Favorite bool `json:"favorite" gorm:"default:true"`
	}

	updatedInfo := new(UpdateStock)
	err = json.NewDecoder(r.Body).Decode(updatedInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't create new stock"))
		return
	}
	stock.Favorite = updatedInfo.Favorite

	DB.Where("symbol=?", symbol).Where("user_refer=?", user_refer).Save(&stock)

	json.NewEncoder(w).Encode(stock)
}

func GetFavorites(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_refer := params["user_refer"]
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	wd := uint(u64)
	var stocks []Stock
	DB.Where("user_refer=?", wd).Where("favorite=?", true).Find(&stocks)
	json.NewEncoder(w).Encode(stocks)
}
