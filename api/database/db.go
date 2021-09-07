package database

import (
	"api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:root@tcp(db:3306)/test"

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could no connect database")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
