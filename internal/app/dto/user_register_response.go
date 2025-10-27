package dto

import (
	"time"

	"github.com/ridhalfajri/pos-api/internal/repository/model"
)

type UserRegisterResponse struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUserRegisterResponse(user model.User, token string) UserRegisterResponse {
	return UserRegisterResponse{
		Id:        user.Id,
		Username:  user.Username,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
