package stocks

import (
	"github.com/gofiber/fiber/v2"
)

func GetStocks(c *fiber.Ctx) error {
	return c.SendString("return all stocks")

}

func GetStock(c *fiber.Ctx) error {
	return c.SendString("return a single stock")

}

func AddStock(c *fiber.Ctx) error {
	return c.SendString("Adds a new stock")

}

func DeleteStock(c *fiber.Ctx) error {
	return c.SendString("Deletes a single stock")

}
