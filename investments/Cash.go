package investments

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Cash struct {
	Currency  string `json:"currency"`
	Amount    uint   `json:"amount" gorm:"default:0"`
	UserRefer uint   `json:"userRefer"`
	//User      User   `gorm:"foreignKey:UserRefer"`
}

func MigrateCash() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to database")
	}
	DB.AutoMigrate(&Cash{})
}

func GetCashInvestments(c *fiber.Ctx) error {

	var user_refer = c.Params("user_refer")
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var money []Cash
	DB.Find(&money, "user_refer=?", wd)
	return c.JSON(&money)
}

func SaveCash(c *fiber.Ctx) error {
	cash := new(Cash)
	if err := c.BodyParser(cash); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&cash)

	return c.JSON(&cash)

}

func DeleteCash(c *fiber.Ctx) error {
	var user_refer = c.Params("user_refer")
	var currency = c.Params("currency")
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert id to uint

	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var cash Cash
	DB.Where("user_refer=?", wd).Where("currency=?", currency).Unscoped().Delete(&cash)
	return c.SendString("Cash asset removed")
}

func UpdateCash(c *fiber.Ctx) error {
	userRefer := c.Params("user_refer")
	currency := c.Params("currency")

	type cashUpdate struct {
		Amount uint `json:"amount"`
	}
	var update cashUpdate
	if err := c.BodyParser(&update); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	result := DB.Model(&Cash{}).
		Where("user_refer = ? AND currency = ?", userRefer, currency).
		Update("amount", update.Amount)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("cash not found")
	}

	var cash Cash
	DB.Where("user_refer = ? AND currency = ?", userRefer, currency).
		First(&cash)
	return c.JSON(&cash)
}
