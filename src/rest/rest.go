package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {

	r := gin.Default()

	r.GET("/mweets", getAllMweeter)
	r.POST("/mweets", createMweet)
	r.PUT("/mweets", updateMweet)
	r.DELETE("/mweets/:id", deleteMweet)
	r.GET("/mweets/:id", getMweeterById)

	r.POST("/users/signin", signInUser)
	r.POST("/users", createUser)
	r.POST("/users/:id/signout", signOutUser)
	r.PUT("/mweets/:id", updateProfile)

	return nil
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

func createUser(ctx *gin.Context) {
	// ...
}

func signInUser(ctx *gin.Context) {
	// ...
}

func signOutUser(ctx *gin.Context) {
	// ...
}

func updateProfile(ctx *gin.Context) {
	// ...
}
