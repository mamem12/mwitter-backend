package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI() *gin.Engine {

	r := gin.Default()

	h, _ := NewHandler()

	r.POST("/users", h.CreateUser)
	r.POST("/users/signin", h.SignInUser)
	r.PUT("/users/:id", h.UpdateProfile)

	return r
}
