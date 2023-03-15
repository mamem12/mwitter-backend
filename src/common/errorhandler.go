package common

import (
	"github.com/gin-gonic/gin"
)

func HandleErrorWithResponse(msg string, httpCode int, ctx *gin.Context) {
	ctx.JSON(httpCode, msg)
}
