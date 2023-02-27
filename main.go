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

}

func main() {

	investments.InitialMigration()
	app := fiber.New()

	app.Get("/", hello)
	Routers(app)

	app.Listen(":3000")

}
