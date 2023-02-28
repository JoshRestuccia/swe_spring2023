package investments

import (
	"fmt"
	"strconv"

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
	//initialize stock table

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	DB.AutoMigrate(&Stock{})
}

func FindUser(id uint, user *User) User {
	//finds the user who has the stock we are trying to add

	DB.Find(&user, "id =?", id)

	if user.ID == id {
		return *user
	}
	return User{}
}

func GetStocks(c *fiber.Ctx) error {

	//returns all stocks of a given user

	var user_refer = c.Params("user_refer")
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stocks []Stock
	DB.Find(&stocks, "user_refer=?", wd)
	for i := range stocks {
		stocks[i].User = FindUser(wd, &stocks[i].User)

	}
	return c.JSON(&stocks)

}

func SaveStock(c *fiber.Ctx) error {

	//adds a new stock
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

	//removes a single stock from a user's portfolio

	var user_refer = c.Params("user_refer")
	var symbol = c.Params("symbol")

	u64, err := strconv.ParseUint(user_refer, 10, 32)
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stock Stock
	DB.Where("user_refer=?", wd).Where("symbol=?", symbol).Delete(&stock)

	//DB.Delete(&stock).Where("user_refer=?", wd).Where("symbol=?", symbol).Find(&stock)
	return c.SendString("Stock deleted")

}

func DeleteStocks(c *fiber.Ctx) error {

	//removes a single stock from a user's portfolio

	var user_refer = c.Params("user_refer")

	u64, err := strconv.ParseUint(user_refer, 10, 32)
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stock Stock
	DB.Where("user_refer=?", wd).Delete(&stock)

	//DB.Delete(&stock).Where("user_refer=?", wd).Where("symbol=?", symbol).Find(&stock)
	return c.SendString("Stock deleted")

}

func UpdateStock(c *fiber.Ctx) error {

	//updates a stock

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
