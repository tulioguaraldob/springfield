package repository

import (
	"github.com/TulioGuaraldoB/springfield/internal/core/domain/model"
	"github.com/TulioGuaraldoB/springfield/internal/core/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserByID(userId uuid.UUID) (*model.User, error) {
	user := new(model.User)
	if err := r.db.Where("user_id = ?", &userId).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
