package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/stocks", allStocks).Methods("GET")

	myRouter.HandleFunc("/users", GetUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", GetUser).Methods("GET")
	myRouter.HandleFunc("/users", CreateUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")

	myRouter.HandleFunc("/stocks/{id}", AddStock).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM")

	InitialMigration()
	handleRequests()
}
