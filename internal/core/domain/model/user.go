package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"size:36"`
	Name      string    `gorm:"not null; size:150"`
	Login     string    `gorm:"not null; size:75"`
	Password  string    `gorm:"not null; size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
