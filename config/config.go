package config

import (
	"gobank/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewMySQLDB() (*gorm.DB, error) {
	var err error
	var DB *gorm.DB
	dsn := "user:password@tcp(127.0.0.1:3306)/people?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to database")

	}

	err = DB.AutoMigrate(&model.Account{})
	return DB, err
}
