package investments

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Stock struct {
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	UserRefer uint    `json:"userRefer"`
	User      User    `gorm:"foreignKey:UserRefer"`
}

func MigrateStocks() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	DB.AutoMigrate(&Stock{})
}

func FindUser(id uint, user *User) User {
	DB.Find(&user, "id =?", id)

	if user.ID == id {
		return *user
	}
	return User{}
}

func GetStocks(c *fiber.Ctx) error {
	var stocks []Stock
	DB.Find(&stocks)
	return c.JSON(&stocks)

}

func GetStock(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	user := c.Params("userRefer")
	var stock Stock

	DB.Find(&stock, user, symbol)

	return c.JSON(&stock)

}

func SaveStock(c *fiber.Ctx) error {
	stock := new(Stock)
	if err := c.BodyParser(stock); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var user User

	DB.Create(&stock)
	stock.User = FindUser(stock.UserRefer, &user)

	return c.JSON(&stock)

}

func DeleteStock(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	user := c.Params("userRefer")
	var stock Stock
	DB.First(&stock, user, symbol)
	if stock.Symbol == "" {
		return c.Status(500).SendString("Stock not found")
	}

	DB.Delete(&stock)
	return c.SendString("Stock deleted")

}

func UpdateStock(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	user := c.Params("userRefer")
	stock := new(Stock)
	DB.First(&stock, symbol, user)
	if stock.Symbol == "" {
		return c.Status(500).SendString("Stock not found")

	}

	if err := c.BodyParser(stock); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&stock)

	return c.JSON(&stock)

}
