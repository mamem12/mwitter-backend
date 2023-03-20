package rest

import (
	"mwitter-backend/src/auth"
	"mwitter-backend/src/common"
	"mwitter-backend/src/rest/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAPI() *gin.Engine {

	router := gin.New()

	router.Use(gin.Recovery())

	UsersRouter(router)
	BookRouter(router)

	return router
}

func UsersRouter(router *gin.Engine) {
	handler, _ := handler.NewUserHandler()

	usersRouterGroup := router.Group("/users")

	usersRouterGroup.POST("/", handler.CreateUser)
	usersRouterGroup.POST("/signin", handler.SignInUser)
	usersRouterGroup.PUT("/:id", Certification(), handler.UpdateProfile)

}

func BookRouter(router *gin.Engine) {

	handler, _ := handler.NewBookHandler()

	booksRouterGroup := router.Group("/books", Certification())

	booksRouterGroup.GET("", handler.GetBookList)
	booksRouterGroup.GET("/:id", handler.GetBookInfoAll)
	booksRouterGroup.GET("/rank/:id", handler.GetBookRank)

}

func Certification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt := auth.JWTToken{}

		err := jwt.TokenValid(ctx)

		if err != nil {
			common.HandleErrorWithResponse(err.Error(), http.StatusInternalServerError, ctx)
			return
		} else {
			ctx.Next()
		}
	}
}
