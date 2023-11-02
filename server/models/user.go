package models

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID           uint  `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"column:name"`
	Email        string `json:"email" gorm:"column:email;unique"`
	Password     []byte	 `json:"-"`
}