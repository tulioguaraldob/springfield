package api

import (
	"github.com/TulioGuaraldoB/springfield/internal/adapters/repository"
	"github.com/TulioGuaraldoB/springfield/internal/adapters/repository/mysql"
	"github.com/TulioGuaraldoB/springfield/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

func GetRoutes() *fiber.App {
	// Adapters
	db := mysql.OpenConnection()

	// Repositories
	userRepository := repository.NewUserRepository(db)

	// Services
	userService := service.NewUserService(userRepository)

	// Controllers
	userController := NewUserController(userService)

	app := fiber.New()
	api := app.Group("api")
	user := api.Group("user")
	user.Get("", userController.GetAllUsers)

	return app
}
