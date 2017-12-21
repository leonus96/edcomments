package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username        string    `json:"username" gorm:"not null; unique"`
	Email           string    `json:"email" gorm:"not null; unique"`
	Fullname        string    `json:"fulname" gorm:"not null"`
	Password        string    `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string    `json:"confirmPassword;om,tempty" gorm:"-"`
	Picture         string    `json:"picture"`
	Comments        []Comment `json:"comments, omitempty"`
}
