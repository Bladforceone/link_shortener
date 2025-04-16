package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go_pro_api/internal/link"
	"go_pro_api/internal/stat"
	"go_pro_api/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
	fmt.Println("Migrations completed")
}
