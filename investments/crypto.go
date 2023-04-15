package investments

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type crypto struct {
	Name        string `json:"name"`
	Amount      uint   `json:"amount" gorm:"default:0"`
	DollarValue uint   `json:"dollar_"`
	UserRefer   uint   `json:"userRefer"`
	//User      User   `gorm:"foreignKey:UserRefer"`
}

func MigrateCrypto() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to database")
	}
	DB.AutoMigrate(&Cash{})
}

func GetCryptoInvestments(c *fiber.Ctx) error {

	var user_refer = c.Params("user_refer")
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert to uint
	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var crypto []crypto
	DB.Find(&crypto, "user_refer=?", wd)
	return c.JSON(&crypto)
}

func GetSingleCrypto(c *fiber.Ctx) error {
	var user_refer = c.Params("user_refer")
	var name = c.Params("name")
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert id to uint

	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var crypto crypto
	DB.Where("user_refer=?", wd).Where("name=?", name).Find(&crypto)
	return c.JSON(&crypto)
}

func SaveCrypto(c *fiber.Ctx) error {
	crypto := new(crypto)
	if err := c.BodyParser(crypto); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&crypto)

	return c.JSON(&crypto)

}

func DeleteCrypto(c *fiber.Ctx) error {
	var user_refer = c.Params("user_refer")
	var name = c.Params("name")
	u64, err := strconv.ParseUint(user_refer, 10, 32)
	//convert id to uint

	if err != nil {
		fmt.Println(err.Error())

	}
	wd := uint(u64)
	var crypto crypto
	DB.Where("user_refer=?", wd).Where("name=?", name).Unscoped().Delete(&crypto)
	return c.SendString("Crypto asset removed")
}

func UpdateCrypto(c *fiber.Ctx) error {
	userRefer := c.Params("user_refer")
	name := c.Params("name")

	type cryptoUpdate struct {
		Amount uint `json:"amount"`
	}
	var update cryptoUpdate
	if err := c.BodyParser(&update); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	result := DB.Model(&crypto{}).
		Where("user_refer = ? AND name = ?", userRefer, name).
		Update("amount", update.Amount)
	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("crypto not found")
	}

	var crypto crypto
	DB.Where("user_refer = ? AND name = ?", userRefer, name).
		First(&crypto)
	return c.JSON(&crypto)
}
