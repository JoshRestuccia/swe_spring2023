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

	app.Get("/stocks", investments.GetStocks)
	app.Get("/stocks/:symbol", investments.GetStock)
	app.Post("/stock", investments.SaveStock)
	app.Delete("/stock/:symbol", investments.DeleteStock)
	app.Put("/stock/:symbol", investments.UpdateStock)

}

func main() {

	investments.InitialMigration()
	investments.MigrateStocks()
	app := fiber.New()

	app.Get("/", hello)
	Routers(app)

	app.Listen(":3000")

}
