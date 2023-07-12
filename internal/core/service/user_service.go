package service

import (
	"github.com/TulioGuaraldoB/springfield/internal/core/domain/model"
	"github.com/TulioGuaraldoB/springfield/internal/core/ports"
	"github.com/TulioGuaraldoB/springfield/pkg/encrypt"
	"github.com/TulioGuaraldoB/springfield/pkg/token"
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

func (s *userService) CreateUser(user *model.User) error {
	user.HashPassword()
	return s.repository.CreateUser(user)
}

func (s *userService) GetUserByCredentials(login, password string) (*string, error) {
	user, err := s.repository.GetUserByCredentials(login, encrypt.ToHash256(password))
	if err != nil {
		return nil, err
	}

	token, err := token.GenerateToken(user.UserID.String(), user.Name)
	if err != nil {
		return nil, nil
	}

	return &token, nil
}
