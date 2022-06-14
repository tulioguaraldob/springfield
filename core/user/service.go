package user

import "github.com/TulioGuaraldoB/springfield/model"

type interfaceService interface {
	getAll() ([]model.User, error)
	show(id uint) (*model.User, error)
	create(user *model.User) error
	delete(id uint) error
}

type service struct {
	repository interfaceRepository
}

func NewService(repository interfaceRepository) interfaceService {
	return &service{
		repository: repository,
	}
}

func (s *service) getAll() ([]model.User, error) {
	return s.repository.getAll()
}

func (s *service) show(id uint) (*model.User, error) {
	return s.repository.show(id)
}

func (s *service) create(user *model.User) error {
	return s.repository.create(user)
}

func (s *service) delete(id uint) error {
	return s.repository.delete(id)
}