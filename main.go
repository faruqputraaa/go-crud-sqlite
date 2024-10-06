package main

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrasi schema Student ke database
	db.AutoMigrate(&Student{})
}

func main() {
	// Inisialisasi database
	initDB()

	e := echo.New()
	Route(e)
	e.Start(":8080")
}
