package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Code  string
	Price uint
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/tryDemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})
	insertProduct := &Product{Code: "D42", Price: 100}
	db.Create(insertProduct)
	fmt.Println(insertProduct.ID, insertProduct.Code, insertProduct.Price)
	readProduct := &Product{}
	db.First(&readProduct, "code = ?", "D42")
	fmt.Println(readProduct.ID, readProduct.Code, readProduct.Price)
	fmt.Println(time.Time{})
}
