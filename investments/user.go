package investments

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3306)/portfullio?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	//Create User structure

	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//Stocksref []*string `json:"stocksref"`
	// Stocks    []Stock  `gorm:"foreignKey:StockRef;constraint:OnDelete:CASCADE,OnDelete:SET NULL;"`
}

//var stocks[] Stock

func InitialMigration() {
	//Create user table
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	DB.AutoMigrate(&User{})

}

func GetUsers(c *fiber.Ctx) error {

	// return all users

	var users []User
	DB.Find(&users)
	return c.JSON(&users)

}

func GetUsersStocks(c *fiber.Ctx) error {
	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stocks []Stock
	var names []string

	//find all stocks matching the user id
	DB.Find(&stocks, "user_refer=?", wd)
	for i := range stocks {
		names = append(names, stocks[i].Name)
	}
	return c.JSON(&names)

}

func GetUsersTotalcrypto(c *fiber.Ctx) error {
	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var crypto []crypto
	var total float64
	DB.Find(&crypto, "user_refer=?", wd)
	for i := range crypto {
		total += float64(crypto[i].Amount * uint(crypto[i].DollarConvert))
	}
	fmt.Printf("Total crypto assets value: $%.2f\n", total)
	return c.JSON(&total)

}

func GetUsersTotalStocks(c *fiber.Ctx) error {
	//returns total value of all stocks owned by the user

	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var stocks []Stock
	var total float64

	//find all stocks matching the user id
	DB.Find(&stocks, "user_refer=?", wd)
	for i := range stocks {
		total += float64(stocks[i].Quantity) * stocks[i].Price

	}

	fmt.Printf("Total stock portfolio value: $%.2f\n", total)

	return c.JSON(&total)
}

func GetUsersTotalCash(c *fiber.Ctx) error {
	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var cash []Cash
	var total float64
	DB.Find(&cash, "user_refer=?", wd)
	for i := range cash {
		total += float64(cash[i].Amount)
	}
	fmt.Printf("Total cash assets value: $%.2f\n", total)
	return c.JSON(&total)
}

func GetUsersTotal(c *fiber.Ctx) error {
	var total float64
	id := c.Params("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var cash []Cash

	DB.Find(&cash, "user_refer=?", wd)
	for i := range cash {
		total += float64(cash[i].Amount)
	}
	fmt.Printf("Total cash assets value: $%.2f\n", total)
	var total2 float64
	var stocks []Stock
	DB.Find(&stocks, "user_refer=?", wd)
	for i := range stocks {
		total2 += float64(stocks[i].Quantity) * stocks[i].Price

	}
	fmt.Printf("Total stock portfolio value: $%.2f\n", total2)
	total += total2
	var total3 float64
	var crypto []crypto
	DB.Find(&crypto, "user_refer=?", wd)
	for i := range crypto {
		total3 += float64(crypto[i].Amount) * crypto[i].DollarConvert

	}
	fmt.Printf("Total crypto portfolio value: $%.2f\n", total3)
	total += total3
	fmt.Printf("Total portfolio value: $%.2f\n", total)
	return c.JSON(&total)

}

func GetUser(c *fiber.Ctx) error {
	//return single user

	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)

}

func SaveUser(c *fiber.Ctx) error {
	//add a new user

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&user)
	return c.JSON(&user)

}

func DeleteUser(c *fiber.Ctx) error {
	//remove a user

	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not found")
	}

	DB.Delete(&user)
	return c.SendString("User is deleted")

}

func UpdateUser(c *fiber.Ctx) error {

	//update a user

	id := c.Params("id")

	user := new(User)
	DB.First(&user, id)
	if user.Username == "" {
		return c.Status(500).SendString("User not found")
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&user)

	return c.JSON(&user)

}
