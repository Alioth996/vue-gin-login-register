package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20); not null" json:"username"`
	Nickname string `gorm:"type:varchar(20);" json:"nickname"`
	Password string `gorm:"type:varchar(100); not null" json:"password"`
	State    int    `gorm:"type:int;default:0" json:"state" `
	Phone    string `gorm:"type:varchar(20);" json:"phone" `
	Email    string `gorm:"type:varchar(40);" json:"email" `
	Address  string `gorm:"type:varchar(200);" json:"address" `
}
