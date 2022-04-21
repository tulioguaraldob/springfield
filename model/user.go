package model

type User struct {
	ID       uint   `gorm:"primaryKey;column:ID;"`
	Name     string `gorm:"column:Name;"`
	Login    string `gorm:"column:Login;"`
	Password string `gorm:"column:Password;"`
}
