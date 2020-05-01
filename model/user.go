package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User Model Defination
type User struct {
	gorm.Model
	Name         string
	Age          string
	Email        string `gorm:"type:varchar(100)"`
	Role         string `gorm:"size:255"`        // set field size to 255
	MemberNumber string `gorm:"unique;not null"` // set member number to unique and not null
}

// TableName Disable table name's pluralization, if set to true, `User`'s table name will be `user`
//db.SingularTable(true)
func (user *User) TableName() string {
	return "user"
}

// ToString show model ditail
func (user User) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s\nemail: %s\nmembernumber: %s\nrole: %s\ncreated: %s", user.ID, user.Name, user.Email, user.MemberNumber, user.Role, user.CreatedAt.String())
}
