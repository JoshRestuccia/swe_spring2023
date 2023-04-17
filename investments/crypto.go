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

type crypto struct {
	Name          string  `json:"name"`
	Amount        uint    `json:"amount" gorm:"default:0"`
	DollarConvert float64 `json:"dollar_" gorm:"default:1.00"`
	UserRefer     uint    `json:"userRefer"`
}

func MigrateCrypto() {
	var err error
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to database")
	}
	DB.AutoMigrate(&crypto{})
}

func GetCryptoInvestments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer, err := strconv.ParseUint(params["user_refer"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	wd := uint(userRefer)
	var cryptoList []crypto
	DB.Find(&cryptoList, "user_refer=?", wd)
	json.NewEncoder(w).Encode(cryptoList)
}

func GetSingleCrypto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer, err := strconv.ParseUint(params["user_refer"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	wd := uint(userRefer)
	var crypto crypto
	DB.Where("user_refer=?", wd).Where("name=?", params["name"]).Find(&crypto)
	json.NewEncoder(w).Encode(crypto)
}

func SaveCrypto(w http.ResponseWriter, r *http.Request) {
	var crypto crypto
	err := json.NewDecoder(r.Body).Decode(&crypto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	DB.Create(&crypto)
	json.NewEncoder(w).Encode(crypto)
}

func DeleteCrypto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer, err := strconv.ParseUint(params["user_refer"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	wd := uint(userRefer)
	var crypto crypto
	DB.Where("user_refer=?", wd).Where("name=?", params["name"]).Unscoped().Delete(&crypto)
	w.Write([]byte("Crypto asset removed"))
}

func UpdateCrypto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userRefer := params["user_refer"]
	name := params["name"]

	type cryptoUpdate struct {
		Amount uint `json:"amount"`
	}
	var update cryptoUpdate
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := DB.Model(&crypto{}).
		Where("user_refer = ? AND name = ?", userRefer, name).
		Update("amount", update.Amount)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "crypto not found", http.StatusNotFound)
		return
	}

	var crypto crypto
	DB.Where("user_refer = ? AND name = ?", userRefer, name).First(&crypto)
	json.NewEncoder(w).Encode(crypto)
}
