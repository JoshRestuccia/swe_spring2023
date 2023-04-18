package main

import (
	"github.com/JoshRestuccia/swe_spring2023/investments"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")

}

func Routers(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", investments.GetUsers)
	api.Get("/users/:id", investments.GetUser)
	api.Get("/users/:id/stocks", investments.GetUsersStocks)
	api.Get("users/:id/totalstocks", investments.GetUsersTotalStocks)
	api.Get("users/:id/totalcash", investments.GetUsersTotalCash)
	api.Get("users/:id/totalcrypto", investments.GetUsersTotalcrypto)
	api.Get("users/:id/total", investments.GetUsersTotal)

	api.Post("/users", investments.SaveUser)
	api.Delete("/users/:id", investments.DeleteUser)
	api.Put("/users/:id", investments.UpdateUser)
	//End of user routes

	api.Get("/stocks/:user_refer", investments.GetStocks)
	api.Get("/stocks/:user_refer/:symbol", investments.GetStock)
	api.Post("/stocks/:user_refer", investments.SaveStock) //changed from "/stocks" to "stocks/:user_refer since it seems like we are storing stock-user pairs and we need the user_refer for that"
	api.Delete("/stocks/:user_refer/:symbol", investments.DeleteStock)
	api.Delete("/stocks/:user_refer", investments.DeleteStocks)
	api.Put("/stocks/:user_refer/:symbol", investments.UpdateStock)
	api.Put("/favorites/:user_refer/:symbol", investments.FavoriteStock)
	api.Get("/favorites/:user_refer", investments.GetFavorites)
	//End of stock routes

	api.Get("/cash/:user_refer", investments.GetCashInvestments)
	api.Post("/cash/:user_refer", investments.SaveCash)
	api.Put("/cash/:user_refer/:currency", investments.UpdateCash)
	api.Delete("/cash/:user_refer/:currency", investments.DeleteCash)
	api.Get("/cash/:user_refer/:currency", investments.GetSingleCash)

	api.Get("/crypto/:user_refer", investments.GetCryptoInvestments)
	api.Post("/crypto/:user_refer", investments.SaveCrypto)
	api.Put("/crypto/:user_refer/:name", investments.UpdateCrypto)
	api.Delete("/crypto/:user_refer/:name", investments.DeleteCrypto)
	api.Get("/crypto/:user_refer/:name", investments.GetSingleCrypto)
}

func main() {
	//Initialize databse tables
	investments.InitialMigration()
	investments.MigrateStocks()
	investments.MigrateCash()
	investments.MigrateCrypto()
	//Initialize Fiber

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", hello)
	Routers(app)

	app.Listen("0.0.0.0:3000")

}
