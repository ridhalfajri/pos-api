package util

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}
type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func SuccessResponse(ctx *gin.Context, httpStatus int, message string, data interface{}) {
	response := BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
		Errors:  nil}
	ctx.JSON(httpStatus, response)
}
func ErrorResponse(ctx *gin.Context, httpStatus int, message string, errorsPayload interface{}) {
	response := BaseResponse{
		Success: false,
		Message: message,
		Data:    nil,
		Errors:  errorsPayload,
	}
	ctx.JSON(httpStatus, response)

}
