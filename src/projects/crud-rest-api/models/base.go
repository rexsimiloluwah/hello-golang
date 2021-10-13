package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// Global db variable
var db *gorm.DB

// The `init()` function runs at the initial start of the program in Go
func init() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables from .env")
	}

	DB_PORT := os.Getenv("DB_PORT")
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	//DB_DRIVER := os.Getenv("DB_DRIVER")

	conn, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DB_USER,
			DB_PASSWORD,
			DB_HOST,
			DB_PORT,
			DB_NAME,
		),
	)
	if err != nil {
		log.Fatal("Error", err.Error())
	}

	db = conn
	// Migrations
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Post{})
	db.Model(&Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

}

// Function for calling the database
func Db() *gorm.DB {
	return db
}
