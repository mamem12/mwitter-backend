package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT interface {
	CreateJWT(userId uint) (string, error)
	TokenValid(ctx *gin.Context) error
}

type JWTToken struct {
}

func (j *JWTToken) CreateJWT(userId uint) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userId"] = userId
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *JWTToken) TokenValid(ctx *gin.Context) error {

	token, err := verifyJWT(ctx)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return fmt.Errorf("")
	}
	return nil
}

func verifyJWT(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := extractJWT(ctx)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("ACCESS_SECRET"), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func extractJWT(ctx *gin.Context) string {

	const BEARER_SCHEMA = "Bearer "
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]

	return tokenString
}
