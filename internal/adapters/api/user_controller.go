package api

import (
	"net/http"

	"github.com/TulioGuaraldoB/springfield/internal/core/domain/dto"
	"github.com/TulioGuaraldoB/springfield/internal/core/ports"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service ports.UserService
}

func NewUserController(service ports.UserService) ports.UserController {
	return &userController{service: service}
}

func (c *userController) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := c.service.GetAllUsers()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	if len(users) == 0 {
		return ctx.Status(http.StatusNotFound).
			JSON(fiber.Map{"error": gorm.ErrRecordNotFound.Error()})
	}

	usersResponse := make([]dto.UserResponse, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, *dto.ParseUserResponse(&user))
	}

	return ctx.Status(http.StatusOK).
		JSON(usersResponse)
}
