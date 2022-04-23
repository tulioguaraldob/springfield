package routes

import (
	"github.com/TulioGuaraldoB/springfield/core/user"
	"github.com/TulioGuaraldoB/springfield/db"
	"github.com/TulioGuaraldoB/springfield/server/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes() *gin.Engine {
	router := gin.Default()

	userRepository := user.NewUserRepository(db.ConnectDB())
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/register", userController.Register)
			v1.POST("/login", userController.Login)
			v1.GET("/user", middlewares.AuthMiddleware(), userController.GetAllUsers)
			v1.GET("/user/:id", middlewares.AuthMiddleware(), userController.GetUserById)
		}
	}

	return router
}
