package user

// import (
// 	"github.com/TulioGuaraldoB/springfield/model"
// 	"github.com/TulioGuaraldoB/springfield/utils/token"
// )

// type interfaceService interface {
// 	getAll() ([]model.User, error)
// 	show(id uint) (*model.User, error)
// 	create(user *model.User) error
// 	delete(id uint) error
// 	getByCredentials(username, password string) (*string, error)
// }

// type service struct {
// 	repository interfaceRepository
// }

// func NewService(repository interfaceRepository) interfaceService {
// 	return &service{
// 		repository: repository,
// 	}
// }

// func (s *service) getAll() ([]model.User, error) {
// 	return s.repository.getAll()
// }

// func (s *service) show(id uint) (*model.User, error) {
// 	return s.repository.show(id)
// }

// func (s *service) create(user *model.User) error {
// 	return s.repository.create(user)
// }

// func (s *service) delete(id uint) error {
// 	return s.repository.delete(id)
// }

// func (s *service) getByCredentials(username, password string) (*string, error) {
// 	user, err := s.repository.getByCredentials(username, password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	token, err := token.GenerateToken(int(user.ID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &token, nil
// }
