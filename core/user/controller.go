package user

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/springfield/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service IUserService
}

func NewUserController(service IUserService) UserController {
	return UserController{
		service: service,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	register := RegisterCredentials{}

	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := model.User{
		Name:     register.Name,
		Login:    register.Login,
		Password: register.Password,
	}

	_, err := c.service.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "registration success!"})
}

func (c *UserController) Login(ctx *gin.Context) {
	signIn := SignInCredentials{}

	if err := ctx.ShouldBindJSON(&signIn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Login:    signIn.Login,
		Password: signIn.Password,
	}

	token, err := c.service.SignIn(user.Login, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login or password incorrect"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.GetUserById(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
