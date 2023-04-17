package main

import (
	"log"
	"net/http"

	"github.com/JoshRestuccia/swe_spring2023/investments"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Routers(r *mux.Router) {
	r.HandleFunc("/users", investments.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", investments.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/stocks", investments.GetUsersStocks).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/totalstocks", investments.GetUsersTotalStocks).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/totalcash", investments.GetUsersTotalCash).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/totalcrypto", investments.GetUsersTotalcrypto).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}/total", investments.GetUsersTotal).Methods(http.MethodGet)

	r.HandleFunc("/users", investments.SaveUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", investments.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/users/{id}", investments.UpdateUser).Methods(http.MethodPut)

	r.HandleFunc("/stocks/{user_refer}", investments.GetStocks).Methods(http.MethodGet)
	r.HandleFunc("/stocks/{user_refer}/{symbol}", investments.GetStock).Methods(http.MethodGet)
	r.HandleFunc("/stocks/{user_refer}", investments.SaveStock).Methods(http.MethodPost)
	r.HandleFunc("/stocks/{user_refer}/{symbol}", investments.DeleteStock).Methods(http.MethodDelete)
	r.HandleFunc("/stocks/{user_refer}", investments.DeleteStocks).Methods(http.MethodDelete)
	r.HandleFunc("/stocks/{user_refer}/{symbol}", investments.UpdateStock).Methods(http.MethodPut)
	r.HandleFunc("/favorites/{user_refer}/{symbol}", investments.FavoriteStock).Methods(http.MethodPut)
	r.HandleFunc("/favorites/{user_refer}", investments.GetFavorites).Methods(http.MethodGet)

	r.HandleFunc("/cash/{user_refer}", investments.GetCashInvestments).Methods(http.MethodGet)
	r.HandleFunc("/cash/{user_refer}", investments.SaveCash).Methods(http.MethodPost)
	r.HandleFunc("/cash/{user_refer}/{currency}", investments.UpdateCash).Methods(http.MethodPut)
	r.HandleFunc("/cash/{user_refer}/{currency}", investments.DeleteCash).Methods(http.MethodDelete)
	r.HandleFunc("/cash/{user_refer}/{currency}", investments.GetSingleCash).Methods(http.MethodGet)

	r.HandleFunc("/crypto/{user_refer}", investments.GetCryptoInvestments).Methods(http.MethodGet)
	r.HandleFunc("/crypto/{user_refer}", investments.SaveCrypto).Methods(http.MethodPost)
	r.HandleFunc("/crypto/{user_refer}/{name}", investments.UpdateCrypto).Methods(http.MethodPut)
	r.HandleFunc("/crypto/{user_refer}/{name}", investments.DeleteCrypto).Methods(http.MethodDelete)
	r.HandleFunc("/crypto/{user_refer}/{name}", investments.GetSingleCrypto).Methods(http.MethodGet)
}

func main() {
	//Initialize databse tables
	investments.InitialMigration()
	investments.MigrateStocks()
	investments.MigrateCash()
	investments.MigrateCrypto()

	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", hello).Methods(http.MethodGet)
	Routers(r)

	log.Fatal(http.ListenAndServe(":3000", r))
}
