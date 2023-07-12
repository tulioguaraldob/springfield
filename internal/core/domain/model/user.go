package model

import (
	"time"

	"github.com/TulioGuaraldoB/springfield/pkg/encrypt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"size:36"`
	Name      string    `gorm:"not null; size:150"`
	Login     string    `gorm:"not null; size:75"`
	Password  string    `gorm:"not null; size:64"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) HashPassword() {
	password := u.Password
	hashedPassword := encrypt.ToHash256(password)
	u.Password = hashedPassword
}
