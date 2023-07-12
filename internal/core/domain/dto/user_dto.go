package dto

import (
	"time"

	"github.com/TulioGuaraldoB/springfield/internal/core/domain/model"
	"github.com/google/uuid"
)

type UserResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func ParseUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		UserID:    user.UserID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}

func ParseUserRequest(userRequest *UserRequest) *model.User {
	return &model.User{
		UserID:   uuid.New(),
		Name:     userRequest.Name,
		Login:    userRequest.Login,
		Password: userRequest.Password,
	}
}
