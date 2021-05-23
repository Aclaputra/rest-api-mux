package main

import (
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

	})

	if err!= nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}
}