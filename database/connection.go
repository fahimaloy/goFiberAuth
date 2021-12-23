package database

import (
	"fmt"

	"github.com/fahimaloy/fibergo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DNS = "fahim:32446+fa@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	connection, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	fmt.Println("Database Connected!")
	connection.AutoMigrate(&models.User{})
}
