package rest

import (
	"mwitter-backend/src/config/logger"

	"github.com/gin-gonic/gin"
)

func RunAPI() *gin.Engine {

	router := gin.New()

	logger.LogFactory()

	router.Use(gin.LoggerWithFormatter(logger.LogFormat))
	router.Use(gin.Recovery())

	// go parsebook.ParseRun()

	UsersRouter(router)

	return router
}

func UsersRouter(router *gin.Engine) *gin.RouterGroup {
	handler, _ := NewHandler()

	usersRouterGroup := router.Group("/users")

	usersRouterGroup.POST("/", handler.CreateUser)
	usersRouterGroup.POST("/signin", handler.SignInUser)
	usersRouterGroup.PUT("/:id", Certification(), handler.UpdateProfile)

	return usersRouterGroup
}
