package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null; size:150"`
	Login    string `gorm:"not null; size:75"`
	Password string `gorm:"not null; size:50"`
}
