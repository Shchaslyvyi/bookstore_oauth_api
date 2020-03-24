package access_token

import (
	"strings"
	"time"

	"github.com/shchaslyvyi/bookstore_oauth_api/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken struct defines an access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Validate func validates the Access Token
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token id.")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("Invalid user id.")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("Invalid client id.")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expiration time.")
	}
	return nil
}

// GetNewAccessToken returns the valid access token
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired function returns the validity state of the access token
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
