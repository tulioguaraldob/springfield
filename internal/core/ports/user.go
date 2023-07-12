package ports

import (
	"github.com/TulioGuaraldoB/springfield/internal/core/domain/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(userId uuid.UUID) (*model.User, error)
	CreateUser(user *model.User) error
	GetUserByCredentials(login, password string) (*model.User, error)
}

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(userId uuid.UUID) (*model.User, error)
	CreateUser(user *model.User) error
	GetUserByCredentials(login, password string) (*string, error)
}

type UserController interface {
	GetAllUsers(ctx *fiber.Ctx) error
	GetUserByUserID(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
