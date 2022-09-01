package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20); not null" json:"username" binding:"required"`
	Nickname string `gorm:"type:varchar(20); not null" json:"nicknameickname"`
	Password string `gorm:"type:varchar(100); not null" json:"password" binding:"required"`
	State    int    `gorm:"type:int;default:0" json:"state" `
	Phone    string `gorm:"type:varchar(20); not null" json:"phone" `
	Email    string `gorm:"type:varchar(40); not null" json:"email" `
	Address  string `gorm:"type:varchar(200); not null" json:"address" `
}
