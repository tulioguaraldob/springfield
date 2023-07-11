package ports

import (
	"github.com/TulioGuaraldoB/springfield/internal/core/domain/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(userId uuid.UUID) (*model.User, error)
}

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(userId uuid.UUID) (*model.User, error)
}

type UserController interface {
	GetAllUsers(ctx *fiber.Ctx) error
}
