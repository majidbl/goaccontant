package controller

import (
  "fmt"
  "goaccontant/model"
)

func GetCash(field, value string) ([]model.Cash, error){
  var cash []model.Cash
  db, err := model.GetDB()
	if err != nil {
		return cash,err
	}
	f := fmt.Sprintf("%s = ?", field)
	fmt.Println(f, value)
	db.Where(f, value).Find(&cash)
	fmt.Println(f, value)
	return cash, nil
}

func AddCash(username, value, typeCash string) (bool, error){
  db, err := model.GetDB()
	if err != nil {
		return false,err
	}
	userz, err := GetUser("user_name", username)
	if err != nil {
	  return false, err
	}
  db.Model(&userz).Association("Cash").Append(model.Cash{Amount:value, TypeCash:typeCash})
  return true, nil
}

func DeleteCash(field, value string)(bool,error){
  db, err := model.GetDB()
	if err != nil {
		return false,err
	}
	var cash model.Cash
	cashs, _ := GetCash(field,value)
	cash = cashs[0]
	db.Delete(&cash)
  return true, nil
  
}