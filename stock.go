package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct {
	Ticker string `json:"Ticker"`
	Name string `json:"Name"`
	Value string `json:"Value"`
}

type Stocks []Stock

func allStocks(w http.ResponseWriter, r*http.Request){
	stocks := Stocks{
		Stock{Ticker: "GOOG", Name: "Google", Value: "1000.00"},
	}

	fmt.Println("Endpoint Hit: All Stocks Endpoint")
	json.NewEncoder(w).Encode(stocks)
}

func AddStock(w http.ResponseWriter, r*http.Request){
	fmt.Println("Endpoint Hit: All Stocks Endpoint")
	w.Header().Set("Content-Type", "application/json")
	var stock Stock
	json.NewDecoder(r.Body).Decode(&stock)
	DB.Create(&stock)
	json.NewEncoder(w).Encode(stock)
}