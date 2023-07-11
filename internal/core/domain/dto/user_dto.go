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

func ParseUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		UserID:    user.UserID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
}
