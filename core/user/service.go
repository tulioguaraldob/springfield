package user

import "github.com/TulioGuaraldoB/springfield/model"

type IUserService interface {
	Register(user *model.User) (*model.User, error)
	SignIn(login, password string) (*string, error)
	GetAllUsers() ([]model.User, error)
	GetUserById(uId uint) (*model.User, error)
}

type UserService struct {
	repository IUserRepository
}

func NewUserService(repository IUserRepository) IUserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Register(user *model.User) (*model.User, error) {
	registerUser, err := s.repository.SaveUser(user)

	if err != nil {
		return registerUser, err
	}

	return registerUser, nil
}

func (s *UserService) SignIn(login, password string) (*string, error) {
	token, err := s.repository.ValidateSignIn(login, password)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repository.GetUsers()
}

func (s *UserService) GetUserById(uId uint) (*model.User, error) {
	user, err := s.repository.GetUserById(uId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
