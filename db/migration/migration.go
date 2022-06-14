package migration

import (
	"github.com/TulioGuaraldoB/springfield/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
