package service

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(userId int64) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtServiceImpl struct {
}

func (service JwtServiceImpl) GenerateToken(userId int64) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	secretKey := []byte(os.Getenv("SECRET_KEY"))
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
func (service JwtServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	parse, err := jwt.Parse(token, service.parse)
	if err != nil {
		return nil, err
	}
	return parse, nil
}

func (service JwtServiceImpl) parse(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return []byte(os.Getenv("SECRET_KEY")), nil
}

func NewJwtService() JwtService {
	return &JwtServiceImpl{}
}
