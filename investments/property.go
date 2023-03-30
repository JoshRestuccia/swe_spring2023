package investments

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Property struct {
	Currency  string `json:"currency"`
	Value    uint   `json:"value" gorm:"default:0"`
	UserRefer uint   `json:"userRefer"`
	//User      User   `gorm:"foreignKey:UserRefer"`
}

func MigrateProperty() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to database")
	}
	DB.AutoMigrate(&Property{})
}