package seed

import (
	"log"

	"github.com/TulioGuaraldoB/springfield/internal/core/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	seedUser(db)
}

func seedUser(db *gorm.DB) {
	users := make([]model.User, 0)
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("Failed to seed users. %s", err.Error())
	}

	if len(users) == 0 {
		db.Create(&model.User{
			UserID:   uuid.New(),
			Name:     "Bart",
			Login:    "el_barto",
			Password: "bartman",
		})
	}
}
