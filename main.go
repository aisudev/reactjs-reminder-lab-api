package main

import (
	"fmt"
	"reminder/domain"

	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	delivery "reminder/feature/delivery"
	repository "reminder/feature/repository"
	usecase "reminder/feature/usecase"
)

var DB *gorm.DB

func init() {
	if err := InitDB(); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	group := e.Group("")

	delivery.NewHandler(
		group, usecase.NewTodoUsecase(
			repository.NewTodoRepository(DB),
		),
	)

	e.Start(":5555")
}

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("DB/reminder.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	AutoMigrate()

	fmt.Println("Database connected...")
	return nil
}

func AutoMigrate() {
	DB.AutoMigrate(&domain.Todo{})
}
