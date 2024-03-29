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
	"github.com/tidwall/gjson"
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

	assert.Equal(200, w.Code)
	// resultUser := &models.User{}
	var jwt string
	err = json.NewDecoder(w.Body).Decode(&jwt)

	assert.NoError(err)

	fmt.Println(jwt)

	// assert.Equal("example@email.com", resultUser.Email)
	// assert.Equal("", resultUser.Password)
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

func TestBookList(t *testing.T) {
	assert := assert.New(t)

	router := RunAPI()

	w := httptest.NewRecorder()

	reader := strings.NewReader("")

	req, err := http.NewRequest("GET", "/books?per=20&page=3&sort=price", reader)

	req.Header.Add("Authorization", "Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzkzNjUzMTMsInVzZXJJZCI6MX0.t26PT6BxMAo63O0DXQN143Tu1Km-J6yDv48C94qxxUQ")

	assert.NoError(err)

	router.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)

	var body *[]models.BookInfo

	err = json.NewDecoder(w.Body).Decode(&body)

	fmt.Println(len(*body))

	// for _, v := range *body {
	// 	fmt.Printf("%+v\n", v)
	// }

	assert.NoError(err)

}

func TestBookInfo(t *testing.T) {
	assert := assert.New(t)

	router := RunAPI()

	w := httptest.NewRecorder()

	reader := strings.NewReader("")

	req, err := http.NewRequest("GET", "/books/3", reader)

	req.Header.Add("Authorization", "Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzkzNjUzMTMsInVzZXJJZCI6MX0.t26PT6BxMAo63O0DXQN143Tu1Km-J6yDv48C94qxxUQ")

	assert.NoError(err)

	router.ServeHTTP(w, req)

	var body string

	err = json.NewDecoder(w.Body).Decode(&body)

	assert.NoError(err)

	assert.Equal(200, w.Code)

	gJson := gjson.Parse(body)
	fmt.Println(gJson)
}

func TestBookRank(t *testing.T) {
	assert := assert.New(t)

	router := RunAPI()

	w := httptest.NewRecorder()

	reader := strings.NewReader("")

	req, err := http.NewRequest("GET", "/books/rank/3", reader)

	req.Header.Add("Authorization", "Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzkzNjUzMTMsInVzZXJJZCI6MX0.t26PT6BxMAo63O0DXQN143Tu1Km-J6yDv48C94qxxUQ")

	assert.NoError(err)

	router.ServeHTTP(w, req)

	assert.Equal(w.Code, http.StatusOK)

	var body string

	err = json.NewDecoder(w.Body).Decode(&body)

	assert.NoError(err)

	fmt.Printf("%+v\n", body)

}
