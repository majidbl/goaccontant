package model

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// Dbcheck for test mysql connetion work or not
func Dbcheck() error {

	db, err := gorm.Open("mysql", "majid72bl:5080075066@(localhost)/goaccontant?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {

		err := db.Close()
		return err
	}
	defer db.Close()
	return nil
}

// GetDB retun Db connetion
func GetDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	//dbHost := os.Getenv("db_host")
	dbDriver := os.Getenv("db_driver")
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName+"?charset=utf8mb4&parseTime=True")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitDB initial db and check error ocurred or not
func InitDB() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// Check model `User`'s table exists or not
	tableExist := db.HasTable(&User{})

	if tableExist {
		return nil
	}

	db.CreateTable(&User{})
	defer db.Close()
	fmt.Println("Table Created Successfully!!!!!")
	return nil
}
