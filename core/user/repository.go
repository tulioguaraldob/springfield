package user

// import (
// 	"time"

// 	"gorm.io/gorm"

// 	"github.com/TulioGuaraldoB/springfield/model"
// )

// type interfaceRepository interface {
// 	getAll() ([]model.User, error)
// 	show(id uint) (*model.User, error)
// 	create(user *model.User) error
// 	delete(id uint) error
// 	getByCredentials(username, password string) (*model.User, error)
// }

// type repository struct {
// 	db *gorm.DB
// }

// func NewRepository(db *gorm.DB) interfaceRepository {
// 	return &repository{
// 		db: db,
// 	}
// }

// func (r *repository) getAll() ([]model.User, error) {
// 	users := []model.User{}
// 	if err := r.db.Find(&users).Error; err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func (r *repository) show(id uint) (*model.User, error) {
// 	user := model.User{}
// 	if err := r.db.First(&user, &id).Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (r *repository) create(user *model.User) error {
// 	return r.db.Create(&user).Error
// }

// func (r *repository) delete(id uint) error {
// 	user := model.User{}
// 	return r.db.Where(&user.ID, &id).
// 		Update("deleted_at", time.Now()).
// 		Error
// }

// func (r *repository) getByCredentials(username, password string) (*model.User, error) {
// 	user := model.User{}
// 	if err := r.db.Model(&user).
// 		Where("Login = ?", username, "Password = ?", password).
// 		Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
