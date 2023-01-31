package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectionDatabase(){
	dsn := "root:@tcp(127.0.0.1:3306)/gin_api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	db.AutoMigrate(&Product{})

	DB = db
}