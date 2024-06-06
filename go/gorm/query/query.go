package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID int64
	//不指定默认 且不设置 则为空
	Name string `gorm:"default:'gg'"`
	Age  int64
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/tryDemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	var user User
	db.Debug().First(&user)
}
