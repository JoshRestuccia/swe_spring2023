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
	app.Get("users/:id/total", investments.GetUsersTotal)

	app.Post("/users", investments.SaveUser)
	app.Delete("/users/:id", investments.DeleteUser)
	app.Put("/users/:id", investments.UpdateUser)
	//End of user routes

	app.Get("/stocks/:user_refer", investments.GetStocks)
	app.Post("/stocks/:user_refer", investments.SaveStock) //changed from "/stocks" to "stocks/:user_refer since it seems like we are storing stock-user pairs and we need the user_refer for that"
	app.Delete("/stocks/:user_refer/:symbol", investments.DeleteStock)
	app.Delete("/stocks/:user_refer", investments.DeleteStocks)
	app.Put("/stocks/:user_refer/:symbol", investments.UpdateStock) //TODO: implement update stock
	//^ changed from /stocks/:user_refer/ to /stocks/:user_refer/:symbol to specify what tuple to update
	//End of stock routes

}

func main() {
	//Initialize databse tables
	investments.InitialMigration()
	investments.MigrateStocks()
	//Initialize Fiber

	app := fiber.New()

	app.Get("/", hello)
	Routers(app)

	app.Listen(":3000")

}
