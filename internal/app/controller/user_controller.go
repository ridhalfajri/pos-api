package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhalfajri/pos-api/internal/app/dto"
	"github.com/ridhalfajri/pos-api/internal/app/service"
	"github.com/ridhalfajri/pos-api/pkg/util"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}
type UserControllerImpl struct {
	userService service.UserService
}

func (controller *UserControllerImpl) Register(ctx *gin.Context) {
	request := dto.UserRegisterRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.ErrorResponse(ctx, http.StatusBadRequest, "invalid or corrupted JSON format.", err.Error())
		return
	}
	user, err := controller.userService.Register(request)
	if err != nil {
		util.ErrorResponse(ctx, http.StatusUnprocessableEntity, "The registration data format is invalid.", err.Error())
		return
	}
	token := "Contoh Token"
	response := dto.ToUserRegisterResponse(user, token)
	util.SuccessResponse(ctx, http.StatusCreated, "cashier account has been successfully created.", response)
}

func (controller *UserControllerImpl) Login(ctx *gin.Context) {
	request := dto.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		util.ErrorResponse(ctx, http.StatusBadRequest, "invalid or corrupted JSON format.", err.Error())
		return
	}
	login, err := controller.userService.Login(request)
	if err != nil {
		util.ErrorResponse(ctx, http.StatusUnprocessableEntity, "The login data format is invalid.", err.Error())
		return
	}
	token := "Contoh Token"
	response := dto.ToUserLoginResponse(login, token)
	util.SuccessResponse(ctx, http.StatusCreated, "login successful. welcome back!", response)

}

func NewUserControllerImpl(userService service.UserService) UserController {
	return &UserControllerImpl{userService: userService}
}
