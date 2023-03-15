package rest

import (
	"encoding/json"
	"fmt"
	"mwitter-backend/src/auth"
	"mwitter-backend/src/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	router := RunAPI()

	w := httptest.NewRecorder()

	createData := models.User{Password: "1234", Email: "example@email.com", Nickname: "mamem"}

	data, _ := json.Marshal(createData)

	reader := strings.NewReader(string(data))

	req, err := http.NewRequest("POST", "/users", reader)

	assert.NoError(err)

	router.ServeHTTP(w, req)

	assert.Equal(201, w.Code)
}

func TestSignInUser(t *testing.T) {
	assert := assert.New(t)

	router := RunAPI()

	w := httptest.NewRecorder()

	getUser := models.User{Password: "1234", Email: "example@email.com"}

	data, _ := json.Marshal(getUser)

	reader := strings.NewReader(string(data))

	req, err := http.NewRequest("POST", "/users/signin", reader)

	assert.NoError(err)

	router.ServeHTTP(w, req)

	resultUser := &models.User{}

	err = json.NewDecoder(w.Body).Decode(resultUser)

	assert.NoError(err)

	assert.Equal(200, w.Code)
	assert.Equal("example@email.com", resultUser.Email)
	assert.Equal("", resultUser.Password)
}

func TestUpdateProfile(t *testing.T) {
	assert := assert.New(t)

	router := RunAPI()

	userId := "1"

	w := httptest.NewRecorder()

	updateDate := models.User{Nickname: "mamem2"}

	data, _ := json.Marshal(updateDate)

	reader := strings.NewReader(string(data))

	req, err := http.NewRequest("PUT", "/users/"+userId, reader)

	jwt := auth.JWTToken{}

	tokenString, _ := jwt.CreateJWT(1)

	req.Header.Add("Authorization", "Bearer "+tokenString)
	assert.NoError(err)

	fmt.Println()

	router.ServeHTTP(w, req)

	msg := ""

	err = json.NewDecoder(w.Body).Decode(&msg)

	assert.NoError(err)
	assert.Equal(msg, userId)
}
