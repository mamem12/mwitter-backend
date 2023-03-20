package main

import (
	"mwitter-backend/src/config/logger"
	"mwitter-backend/src/rest"

	"github.com/gin-gonic/gin"
)

func main() {

	r := rest.RunAPI()

	logger.LogFactory()

	r.Use(gin.LoggerWithFormatter(logger.LogFormat))

	r.Run(":8080")
}
