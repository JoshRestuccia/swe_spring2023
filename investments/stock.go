package investments

import (
	"errors"
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
	Quantity  int     `json:"quantity" gorm:"default:1"`
	UserRefer uint    `json:"userRefer"`
	//	User      User    `gorm:"foreignKey:UserRefer"`
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

func CreateResponseStock(stock Stock) Stock {
	return Stock{Symbol: stock.Symbol, Name: stock.Name, Price: stock.Price, Quantity: stock.Quantity, UserRefer: stock.UserRefer}
}

func GetStocks(c *fiber.Ctx) error {

	//returns all stocks of a given user

	var user_refer = c.Params("user_refer")
	//get the user id number as a string
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stocks []Stock
	//find all stocks matching the user id
	DB.Find(&stocks, "user_refer=?", wd)
	// for i := range stocks {
	// 	stocks[i].User = FindUser(wd, &stocks[i].User)

	// }
	return c.JSON(&stocks)

}

func SaveStock(c *fiber.Ctx) error {
	//Should be added to add stock with user_refer
	//adds a new stock
	stock := new(Stock)
	if err := c.BodyParser(stock); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	//var user User

	DB.Create(&stock)
	//stock.User = FindUser(stock.UserRefer, &user)

	return c.JSON(&stock)

}

func DeleteStock(c *fiber.Ctx) error {

	//removes a single stock from a user's portfolio

	var user_refer = c.Params("user_refer")
	var symbol = c.Params("symbol")
	// get the user id number and stock symbol as strings

	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert id to uint

	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stock Stock
	//delete stocks matching the user id and stock symbol
	DB.Where("user_refer=?", wd).Where("symbol=?", symbol).Delete(&stock)

	return c.SendString("Stock deleted")

}

func DeleteStocks(c *fiber.Ctx) error {

	//same logic as DeleteStock() but no symbol parameter

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

func findStock(symbol string, id uint, stock Stock) error {
	DB.Where("symbol=?", symbol).Where("user_refer=?", id).Find(&stock)
	if stock.Symbol == "" {
		return errors.New("Stock not found")
	}
	return nil
}

func ReturnStock(symbol string, id uint, stock Stock) Stock {
	DB.Where("symbol=?", symbol).Where("user_refer=?", id).Find(&stock)
	return stock
}

func UpdateStock(c *fiber.Ctx) error {

	//updates a stock

	symbol := c.Params("symbol")
	user := c.Params("user_refer")
	var stock Stock
	u64, er := strconv.ParseUint(user, 10, 32)
	//convert to uint
	if er != nil {
		fmt.Println(er.Error())

	}
	wd := uint(u64)

	err := findStock(symbol, wd, stock)
	stock = ReturnStock(symbol, wd, stock)
	if err != nil {
		return c.Status(500).SendString("Stock not found")

	}

	// if err := c.BodyParser(stock); err != nil {
	// 	return c.Status(500).SendString("Something went wrong")
	// }
	type UpdateStock struct {
		Symbol   string  `json:"symbol"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity" gorm:"default:1"`
	}

	updatedInfo := new(UpdateStock)
	if err := c.BodyParser(updatedInfo); err != nil {
		return c.Status(500).JSON("Can't create new stock")
	}
	stock.Symbol = updatedInfo.Symbol
	stock.Name = updatedInfo.Name
	stock.Quantity = updatedInfo.Quantity
	stock.Price = updatedInfo.Price

	DB.Where("symbol=?", symbol).Where("user_refer=?", user).Save(&stock)

	//NewStock := CreateResponseStock(stock)

	return c.Status(200).JSON(&stock)

}
