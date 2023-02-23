package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI() *gin.Engine {

	r := gin.Default()

	h, _ := NewHandler()

	r.GET("/mweets", getAllMweeter)
	r.POST("/mweets", createMweet)
	r.PUT("/mweets", updateMweet)
	r.DELETE("/mweets/:id", deleteMweet)
	r.GET("/mweets/:id", getMweeterById)

	r.POST("/users", h.CreateUser)
	r.POST("/users/signin", h.SignInUser)
	// r.POST("/users/:id/signout", signOutUser)
	r.PUT("/users/:id", h.UpdateProfile)

	return r
}

func getAllMweeter(ctx *gin.Context) {
	// ...
}

func createMweet(ctx *gin.Context) {
	// ...
}

func updateMweet(ctx *gin.Context) {
	// ...
}

func deleteMweet(ctx *gin.Context) {
	// ...
}

func getMweeterById(ctx *gin.Context) {
	// ...
}
