package dto

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
