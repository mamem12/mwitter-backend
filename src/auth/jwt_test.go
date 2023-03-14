package auth

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	assert := assert.New(t)
	jwtT := JWTToken{}

	tokenString, err := jwtT.CreateJWT(1)

	assert.NoError(err)
	assert.NotEqual(tokenString, "")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = httptest.NewRequest("", "/", nil)
	ctx.Request.Header.Add("Authorization", "Bearer "+tokenString)

	header := ctx.GetHeader("Authorization")
	fmt.Println(header)
	err = jwtT.TokenValid(ctx)
	assert.NoError(err)
}
