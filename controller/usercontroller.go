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
	errs, unique := uniqueCheck(user.MemberNumber)

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
func GetUser(mn string) model.User {
	db, err := model.GetDB()
	if err != nil {
		panic(err)
	}
	var users []model.User
	db.Where("member_number = ?", mn).First(&users)
	fmt.Println(len(users))
	for _, user := range users {
		fmt.Println(user.Name)

	}
	return users[0]
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

func uniqueCheck(mn string) ([]error, bool) {
	db, _ := model.GetDB()
	var user []model.User
	errors := db.Where("member_number = ?", mn).First(&user).GetErrors()
	if len(errors) > 0 {
		return errors, false
	}

	var users []model.User
	db.Where("member_number = ?", mn).First(&users)
	if len(users) > 0 {
		return nil, false
	}
	return nil, true
}
