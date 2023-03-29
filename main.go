package main

import (
	"github.com/JoshRestuccia/swe_spring2023/investments"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")

}

func Routers(app *fiber.App) {
	app.Get("/users", investments.GetUsers)
	app.Get("/users/:id", investments.GetUser)
	app.Get("/users/:id/stocks", investments.GetUsersStocks)
	app.Get("users/:id/totalStocks", investments.GetUsersTotalStocks)

	app.Post("/users", investments.SaveUser)
	app.Delete("/users/:id", investments.DeleteUser)
	app.Put("/users/:id", investments.UpdateUser)
	//End of user routes

	app.Get("/stocks/:user_refer", investments.GetStocks)
	app.Get("/stocks/:user_refer/:symbol", investments.GetStock)
	app.Post("/stocks/:user_refer", investments.SaveStock) //changed from "/stocks" to "stocks/:user_refer since it seems like we are storing stock-user pairs and we need the user_refer for that"
	app.Delete("/stocks/:user_refer/:symbol", investments.DeleteStock)
	app.Delete("/stocks/:user_refer", investments.DeleteStocks)
	app.Put("/stocks/:user_refer/:symbol", investments.UpdateStock)
	app.Put("/stock/:user_refer/:symbol", investments.FavoriteStock)
	app.Get("/favorites/:user_refer", investments.GetFavorites)
	//End of stock routes

	app.Get("/cash/:user_refer", investments.GetCashInvestments)
	app.Post("/cash/:user_refer", investments.SaveCash)
	app.Put("/cash/:user_refer/:currency", investments.UpdateCash)
	app.Delete("/cash/:user_refer/:currency", investments.DeleteCash)
	app.Get("/cash/:user_refer/:currency", investments.GetSingleCash)
}

func main() {
	//Initialize databse tables
	investments.InitialMigration()
	investments.MigrateStocks()
	investments.MigrateCash()
	//Initialize Fiber

	app := fiber.New()

	app.Get("/", hello)
	Routers(app)

	app.Listen(":3000")

}
