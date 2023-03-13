package rest

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"mwitter-backend/src/dblayer"
	"mwitter-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerInterface interface {
	CreateUser(ctx *gin.Context)
	SignInUser(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
	db, err := dblayer.NewORM("test", gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Handler{
		db: db,
	}, nil
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	if h.db == nil {
		return
	}

	user := &models.User{}
	err := ctx.ShouldBindJSON(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existUser, err := h.db.SelectUserByEmail(user.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	fmt.Printf("%+v\n", existUser)
	fmt.Printf("%s\n", existUser.Email)

	if existUser.Email != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "가입이 중복되었습니다."})
		return
	}

	password := user.Password
	hash := sha256.New()
	_, err = hash.Write([]byte(password))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	md := hash.Sum(nil)

	user.Password = hex.EncodeToString(md)

	err = h.db.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, user.ID)
}

func (h *Handler) SignInUser(ctx *gin.Context) {

	if h.db == nil {
		return
	}

	user := &models.User{}
	err := ctx.ShouldBindJSON(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password := user.Password
	hash := sha256.New()
	_, err = hash.Write([]byte(password))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	md := hash.Sum(nil)
	user.Password = hex.EncodeToString(md)

	userInfo, err := h.db.SignInUser(user.Email, user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userInfo.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
	}

	ctx.JSON(http.StatusOK, userInfo)
}

func (h *Handler) UpdateProfile(ctx *gin.Context) {

	if h.db == nil {
		return
	}

	user := &models.User{}
	err := ctx.ShouldBindJSON(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, isId := ctx.Params.Get("id")

	if !isId {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.UpdateProfile(id, user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, id)
}
