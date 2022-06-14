package user

import (
	"time"

	"github.com/TulioGuaraldoB/springfield/model"
	"gorm.io/gorm"
)

type RegisterCredentials struct {
	Name     string `json:"name"`
	Login    string `json:"username"`
	Password string `json:"password"`
}

type SignInCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uint           `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func UserToResponse(user *model.User, res *UserResponse) {
	*res = UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Name:      user.Name,
	}
}

func RequestToUser(req *UserRequest, user *model.User) {
	*user = model.User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:     req.Name,
		Login:    req.Login,
		Password: req.Password,
	}
}
