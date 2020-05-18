package controller

import (
	"fmt"
	"goaccontant/model"
)

// CreateUser add a user to DB
func CreateUser(user *model.User) (bool, error) {

	db, err := model.GetDB()
	if err != nil {
		return false, nil
	}
	var errs []error
	errs, unique := UniqueCheck(user.Email)

	//fmt.Println("uniqueCheck Answer is ", unique)
	if unique {
		db.Create(&user)
		fmt.Println("user created successfully")
		return true, nil
	}
	if (len(errs)) > 0 {
		return false, err
	}
	return false, nil

}

// GetUser return specific user
func GetUser(field, value string) (model.User, error) {
  var user model.User
	db, err := model.GetDB()
	if err != nil {
		return user, err
	}
	var users []model.User
	q := fmt.Sprintf("%s = ?", field)
	db.Where(q, value).First(&users)
	//fmt.Println(len(users))
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//}
	//TODO: IF USER DID NOT FIND
	if len(users) < 1{
	  errors := db.Where(q, value).First(&user).GetErrors()
	  if len(errors) > 0 {
	    return user, errors[0]
	  }
	  
	}
	return users[0], nil
}

//GetAllUser get full users
func GetAllUser() ([]model.User, error) {
	db, err := model.GetDB()
	if err != nil {
		return nil, err
	}
	var users []model.User
	db.Find(&users)
	return users, nil
}

func UniqueCheck(mn string) ([]error, bool) {
	db, _ := model.GetDB()
	var user []model.User

	var users []model.User
	db.Where("email = ?", mn).First(&users)
	errors := db.Where("email = ?", mn).First(&user).GetErrors()
	if len(users) > 0 {
		return errors, false
	}
	//fmt.Println(users[0])
	//fmt.Println(len(users))
	return errors, true
}
