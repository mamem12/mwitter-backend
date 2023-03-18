package rest

import (
	"mwitter-backend/src/config/logger"
	"mwitter-backend/src/parsebook"

	"github.com/gin-gonic/gin"
)

func RunAPI() *gin.Engine {

	router := gin.New()

	logger.LogFactory()

	router.Use(gin.LoggerWithFormatter(logger.LogFormat))
	router.Use(gin.Recovery())

	go parsebook.ParseRun()

	UsersRouter(router)
	BookRouter(router)

	return router
}

func UsersRouter(router *gin.Engine) {
	handler, _ := NewHandler()

	usersRouterGroup := router.Group("/users")

	usersRouterGroup.POST("/", handler.CreateUser)
	usersRouterGroup.POST("/signin", handler.SignInUser)
	usersRouterGroup.PUT("/:id", Certification(), handler.UpdateProfile)

}

func BookRouter(router *gin.Engine) {

	handler, _ := parsebook.NewBookHandler()

	booksRouterGroup := router.Group("/books")

	booksRouterGroup.GET("/", handler.GetBookList)
	booksRouterGroup.GET("/:id", handler.GetBookInfoAll)

}
