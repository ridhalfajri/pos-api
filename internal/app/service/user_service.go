package service

import (
	"errors"

	"github.com/ridhalfajri/pos-api/internal/app/dto"
	"github.com/ridhalfajri/pos-api/internal/repository"
	"github.com/ridhalfajri/pos-api/internal/repository/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(request dto.UserRegisterRequest) (model.User, error)
	Login(request dto.UserLoginRequest) (model.User, error)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (service *UserServiceImpl) Register(request dto.UserRegisterRequest) (model.User, error) {
	user := model.User{}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Username = request.Username
	user.Password = string(password)
	user.Role = "kasir"
	register, err := service.userRepository.Register(user)
	if err != nil {
		return user, err
	}
	return register, nil
}
func (service *UserServiceImpl) Login(request dto.UserLoginRequest) (model.User, error) {
	username := request.Username

	user, err := service.userRepository.FindByUsername(username)
	if err != nil {
		return user, errors.New("the email or password you entered is incorrect")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return user, errors.New("the email or password you entered is incorrect")
	}
	return user, nil
}
