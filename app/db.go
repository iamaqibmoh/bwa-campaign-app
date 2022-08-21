package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func DBConnect() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
