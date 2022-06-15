package user

import (
	"net/http"
	"strconv"

	"github.com/TulioGuaraldoB/springfield/model"
	"github.com/TulioGuaraldoB/springfield/utils/encrypt"
	"github.com/gin-gonic/gin"
)

type controller struct {
	service interfaceService
}

func NewUserController(service interfaceService) controller {
	return controller{
		service: service,
	}
}

func (c *controller) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.getAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allUsersRes := []UserResponse{}
	userRes := UserResponse{}

	for _, user := range users {
		UserToResponse(&user, &userRes)
		allUsersRes = append(allUsersRes, userRes)
	}

	ctx.JSON(http.StatusOK, allUsersRes)
}

func (c *controller) GetUserById(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.show(uint(userId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRes := UserResponse{}
	UserToResponse(user, &userRes)

	ctx.JSON(http.StatusOK, userRes)
}

func (c *controller) Register(ctx *gin.Context) {
	userInput := UserRequest{}
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := model.User{}
	RequestToUser(&userInput, &user)

	hashedPassword := encrypt.ToHash256(user.Password)
	user.Password = hashedPassword

	if err := c.service.create(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "user inserted successfully!",
		"user":    userInput,
	})
}
