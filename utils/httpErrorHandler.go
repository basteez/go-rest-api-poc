package utils

import (
	"github.com/gin-gonic/gin"
)

type ErrorDto struct {
	message string
}

func HandleHttpError(err error, message string, status int, context *gin.Context) {
	if err != nil {
		LogOnError(err, message)
		context.JSON(status, ErrorDto{message: message})
		context.Abort()
	}
}
