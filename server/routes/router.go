package routes

import (
	"github.com/TulioGuaraldoB/springfield/core/user"
	"github.com/TulioGuaraldoB/springfield/db"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes() *gin.Engine {
	router := gin.Default()

	userRepository := user.NewRepository(db.ConnectDB())
	userService := user.NewService(userRepository)
	userController := user.NewUserController(userService)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.POST("/register", userController.Register)
				user.GET("/", userController.GetAllUsers)
				user.GET("/:id", userController.GetUserById)
			}
		}
	}

	return router
}
