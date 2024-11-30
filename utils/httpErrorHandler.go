package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorDto struct {
	Message string `json:"message"`
}

func HandleHttpError(err error, message string, status int, context *gin.Context) {
	if err != nil {
		LogOnError(err, message)
		context.JSON(status, ErrorDto{Message: message})
	}
}
