package user

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/TulioGuaraldoB/springfield/model"
	"github.com/TulioGuaraldoB/springfield/utils/token"
)

type IUserRepository interface {
	SaveUser(user *model.User) (*model.User, error)
	ValidateSignIn(login, password string) (string, error)
	GetUsers() ([]model.User, error)
	GetUserById(userId uint64) (*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

var user *model.User

func (r *UserRepository) SaveUser(user *model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) ValidateSignIn(login, password string) (string, error) {
	userToSign := model.User{}

	err := r.db.Where("login = ? AND password = ?", login, password).
		Take(&userToSign).
		Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, userToSign.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(int(userToSign.ID))
	if err != nil {
		return "", nil
	}

	return token, nil
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var allUsers []model.User

	if err := r.db.Find(&allUsers).Error; err != nil {
		return allUsers, err
	}

	return allUsers, nil
}

func (r *UserRepository) GetUserById(userId uint64) (*model.User, error) {
	var user *model.User

	if err := r.db.First(&user, &userId).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func BeforeSave() error {
	var u *model.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Login = html.EscapeString(strings.TrimSpace(u.Login))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
