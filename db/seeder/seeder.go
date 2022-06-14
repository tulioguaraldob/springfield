package seeder

import (
	"time"

	"github.com/TulioGuaraldoB/springfield/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.Model(&model.User{}).Create(&model.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     "Homer",
		Login:    "beer",
		Password: "donuts",
	})
}
