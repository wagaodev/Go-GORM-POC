package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	dsn := os.Getenv("DB_CONNECTION")
	fmt.Println(dsn)
}
