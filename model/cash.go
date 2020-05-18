package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Cash Model Defination
type Cash struct {
	gorm.Model
	Amount         string `gorm:"type:varchar(256);not null"`
	TypeCash     string `gorm:"type:varchar(256);not null"`
	UserUserName string
	// set field size to 255
}

// TableName Disable table name's pluralization, if set to true, `Cash`'s table name will be `cash`
//db.SingularTable(true)
func (cash *Cash) TableName() string {
	return "cash"
}

// ToString show model detail
func (cash Cash) ToString() string {
	return fmt.Sprintf("id: %d\namount: %s\ntypecash: %s\ncreated: %s", cash.ID, cash.Amount, cash.TypeCash, cash.CreatedAt.String())
}
