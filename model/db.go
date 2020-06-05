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
	// Check model `User`'s and Cash table exists or not
	userTableExist := db.HasTable(&User{})
	cashTableExist := db.HasTable(&Cash{})

	if !userTableExist {
	  db.CreateTable(&User{})
	  fmt.Println("User Table Created Successfully!!!!!")
	}else{
	  fmt.Println("User Table not Created because is Exist!!!!!")
	}
	if !cashTableExist {
	  db.CreateTable(&Cash{})
	  fmt.Println("Cash Table Created Successfully!!!!!")
	}else{
	  fmt.Println("Cash Table not created because is Exist!!!!!")
	}

	defer db.Close()
	return nil
}

/*func CreateUser(user &User)(bool, error){
  db, err := GetDB()
  if err != nil {
    return false, err
  }
  if err := db.Where("member_number = ?", user.MemberNumber).First(&user).Error; gorm.IsRecordNotFoundError(err) {
  // record not found
  db.Create(user)
  return true, nil
  } 
  return false, err
}*/