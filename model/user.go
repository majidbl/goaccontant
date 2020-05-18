package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User Model Defination
type User struct {
	gorm.Model
	UserName         string
	Password         string `gorm:"type:varchar(256);not null"`
	Email            string `gorm:"type:varchar(100)";not null`
	Role             string `gorm:"size:255"` 
	// set field size to 255
	Cash  []Cash `gorm:"foreignkey:UserUserName;association_foreignkey:UserName"`
}


// TableName Disable table name's pluralization, if set to true, `User`'s table name will be `user`
//db.SingularTable(true)
func (user *User) TableName() string {
	return "user"
}

// ToString show model ditail
func (user User) ToString() string {
	return fmt.Sprintf("id: %d\nusername: %s\nemail: %s\nrole: %s\ncreated: %s", user.ID, user.UserName, user.Email, user.Role, user.CreatedAt.String())
}
