package dto

import "github.com/ridhalfajri/pos-api/internal/repository/model"

type UserLoginResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func ToUserLoginResponse(user model.User, token string) UserLoginResponse {
	return UserLoginResponse{
		Id:       user.Id,
		Username: user.Username,
		Token:    token,
	}
}
