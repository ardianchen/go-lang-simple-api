package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func InitDb() *gorm.DB {
	// check load env
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
		panic("check env file")
	}
	// config
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	db, err := gorm.Open("mysql", username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Asia%2FKolkata")
	// defer db.Close()
	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	// db = conn

	//Printing query
	db.LogMode(true)

	//Automatically create migration as per model
	db.AutoMigrate(
	// &mdl.User{},
	)
	return db
}

// func GetDB() *gorm.DB {
// 	return db
// }

func CloseDb() {
	db.Close()
}
