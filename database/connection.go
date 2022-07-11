package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root@localhost:3306"), &gorm.Config{})

	if err != nil {
		panic("Couldn't establish connection with database")
	}
}
