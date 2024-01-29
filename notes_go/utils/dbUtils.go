package utils

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var url = "root:123456@tcp(127.0.0.1:3306)/notes?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectToDB() {

	var err error
	DB, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to mysql : %v", err)
	}
}
