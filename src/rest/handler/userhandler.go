package handler

import (
	"mwitter-backend/src/auth"
	"mwitter-backend/src/common"
	"mwitter-backend/src/dblayer"
	"mwitter-backend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandlerInterface interface {
	CreateUser(ctx *gin.Context)
	SignInUser(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type UserHandler struct {
	db dblayer.UserLayer
}

func NewUserHandler() (UserHandlerInterface, error) {
	db, err := dblayer.NewORM("test", gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &UserHandler{
		db: db.UserORM,
	}, nil
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {

	if h.db == nil {
		return
	}

	user := &models.User{}
	err := ctx.ShouldBindJSON(user)

	if err != nil {
		common.HandleErrorWithResponse(err.Error(), http.StatusBadRequest, ctx)
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

	if existUser.Email != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "가입이 중복되었습니다."})
		return
	}

	hashStr, err := common.StrToHash(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user.Password = hashStr

	err = h.db.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, user.ID)
}

func (h *UserHandler) SignInUser(ctx *gin.Context) {

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

	hashStr, err := common.StrToHash(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	user.Password = hashStr

	userInfo, err := h.db.SignInUser(user.Email, user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userInfo.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
	}

	jwt := auth.JWTToken{}

	token, _ := jwt.CreateJWT(userInfo.ID)

	ctx.JSON(http.StatusOK, token)
}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {

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
