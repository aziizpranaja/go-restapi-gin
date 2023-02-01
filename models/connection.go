package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectionDatabase(){
	dsn := os.Getenv("DB")
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	db.AutoMigrate(&Product{}, &User{})

	DB = db
}