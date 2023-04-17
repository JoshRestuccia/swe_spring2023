package investments

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Cash struct {
	Currency  string `json:"currency"`
	Amount    uint   `json:"amount" gorm:"default:0"`
	UserRefer uint   `json:"userRefer"`
	//User      User   `gorm:"foreignKey:UserRefer"`
}

func MigrateCash() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to database")
	}
	DB.AutoMigrate(&Cash{})
}

func GetCashInvestments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer, err := strconv.ParseUint(params["user_refer"], 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	wd := uint(userRefer)
	var money []Cash
	DB.Find(&money, "user_refer=?", wd)
	json.NewEncoder(w).Encode(money)
}

func GetSingleCash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer, err := strconv.ParseUint(params["user_refer"], 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	wd := uint(userRefer)
	currency := params["currency"]
	var cash Cash
	DB.Where("user_refer=?", wd).Where("currency=?", currency).Find(&cash)
	json.NewEncoder(w).Encode(cash)
}

func SaveCash(w http.ResponseWriter, r *http.Request) {
	var cash Cash
	err := json.NewDecoder(r.Body).Decode(&cash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	DB.Create(&cash)
	json.NewEncoder(w).Encode(cash)
}

func DeleteCash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer, err := strconv.ParseUint(params["user_refer"], 10, 32)
	if err != nil {
		fmt.Println(err.Error())
	}
	wd := uint(userRefer)
	currency := params["currency"]
	var cash Cash
	DB.Where("user_refer=?", wd).Where("currency=?", currency).Unscoped().Delete(&cash)
	fmt.Fprintf(w, "Cash asset removed")
}

func UpdateCash(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer := params["user_refer"]
	currency := params["currency"]
	type cashUpdate struct {
		Amount uint `json:"amount"`
	}
	var update cashUpdate
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := DB.Model(&Cash{}).
		Where("user_refer = ? AND currency = ?", userRefer, currency).
		Update("amount", update.Amount)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "cash not found", http.StatusNotFound)
		return
	}
	var cash Cash
	DB.Where("user_refer = ? AND currency = ?", userRefer, currency).First(&cash)
	json.NewEncoder(w).Encode(cash)
}
