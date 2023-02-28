package investments

import (
	"fmt"

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

// func FindUser(id int, user * User) error {
// 	DB.Find(&user, "id =?", id)
// 	if user.ID ==0{
// 		return errors.New("User not found")

// 	}
// 	return nil
// }

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
