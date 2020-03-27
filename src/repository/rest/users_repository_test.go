package rest

import (
	"fmt"
	"golang-restclient/rest"
	"net/http"
	"os"
	"testing"

	"github.com/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test Cases.")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromAPI(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid rest-client response whet trying to login user.", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalig login credentials", "status": "404", "error":"not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when trying to login user.", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalig login credentials", "status": 404, "error":"not_found"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalig login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "Guzeppe", "last_name": "Dawidoff", "email": "gd@gmail.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error when tryong to unmarshall users login response.", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://api.bookstore.com/users/login",
		ReqBody:      `{"email":"email@gmail.com", "password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "Guzeppe", "last_name": "Dawidoff", "email": "gd@gmail.com"}`,
	})
	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Dawidoff", user.LastName)
	assert.EqualValues(t, "Guzeppe", user.FirstName)
	assert.EqualValues(t, "gd@gmail.com", user.Email)
}
