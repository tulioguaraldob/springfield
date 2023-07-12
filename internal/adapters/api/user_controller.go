package api

import (
	"net/http"

	"github.com/TulioGuaraldoB/springfield/internal/core/domain/dto"
	"github.com/TulioGuaraldoB/springfield/internal/core/ports"
	"github.com/google/uuid"
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

func (c *userController) GetUserByUserID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := c.service.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).
		JSON(dto.ParseUserResponse(user))
}

func (c *userController) CreateUser(ctx *fiber.Ctx) error {
	userRequest := new(dto.UserRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.service.CreateUser(dto.ParseUserRequest(userRequest)); err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusCreated).
		JSON(fiber.Map{
			"message": "user created successfully!",
			"user":    userRequest,
		})
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	userCredentials := new(dto.Credentials)
	if err := ctx.BodyParser(userCredentials); err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	token, err := c.service.GetUserByCredentials(userCredentials.Login, userCredentials.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(http.StatusNotFound).
				JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(http.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).
		JSON(*token)
}
