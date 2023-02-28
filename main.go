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
	app.Post("/user", investments.SaveUser)
	app.Delete("/user/:id", investments.DeleteUser)
	app.Put("/user/:id", investments.UpdateUser)
	//End of user routes

	app.Get("/stocks/:user_refer", investments.GetStocks)
	app.Post("/stock", investments.SaveStock)
	app.Delete("/stock/:user_refer/:symbol", investments.DeleteStock) //todo: implement delete stock

	app.Put("/stock/:user_refer", investments.UpdateStock) //todo: implement update stock

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
