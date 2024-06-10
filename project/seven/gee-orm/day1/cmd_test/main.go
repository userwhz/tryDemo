package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	geerom "tryDemo/project/seven/gee-orm/day1"
)

func main() {
	engine, _ := geerom.NewEngine("sqlite3", "seven.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
