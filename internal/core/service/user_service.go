package service

import (
	"github.com/TulioGuaraldoB/springfield/internal/core/domain/model"
	"github.com/TulioGuaraldoB/springfield/internal/core/ports"
	"github.com/google/uuid"
)

type userService struct {
	repository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) ports.UserService {
	return &userService{repository: repository}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.repository.GetAllUsers()
}

func (s *userService) GetUserByID(userId uuid.UUID) (*model.User, error) {
	return s.repository.GetUserByID(userId)
}
