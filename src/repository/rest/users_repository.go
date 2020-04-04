package rest

import (
	"encoding/json"
	"golang-restclient/rest"
	"time"

	"github.com/shchaslyvyi/bookstore_oauth_api/src/domain/users"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

// RestUsersRepository interface
type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

// NewRestUserRepository function to access an dbRepository struct
func NewRestUserRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest-client response whet trying to login user.")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to login user.")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("Error when tryong to unmarshall users login response.")
	}
	return &user, nil
}
