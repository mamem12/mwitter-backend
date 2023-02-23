package rest

import (
	"mwitter-backend/src/dblayer"
	"mwitter-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetAllMweeter(ctx *gin.Context)
	CreateMweet(ctx *gin.Context)
	UpdateMweet(ctx *gin.Context)
	DeleteMweet(ctx *gin.Context)
	GetMweeterById(ctx *gin.Context)

	CreateUser(ctx *gin.Context)
	SignInUser(ctx *gin.Context)
	SignOutUser(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (*Handler, error) {
	return new(Handler), nil
}

func (h *Handler) GetAllMweeter(ctx *gin.Context) {
	if h.db == nil {
		return
	}

	mweeters, err := h.db.GetAllMweeter()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mweeters)
}

func (h *Handler) CreateMweet(ctx *gin.Context) {
	if h.db == nil {
		return
	}

	user := &models.User{}
	err := ctx.ShouldBindJSON(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.CreateMweet(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
