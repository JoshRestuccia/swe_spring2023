package investments

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Stock struct {
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	// UserRefer string  `json:"user_email"`
	// User      User    `gorm:"foreignKey:UserRefer"`
}

func MigrateStocks() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	DB.AutoMigrate(&Stock{})
}
