package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/shchaslyvyi/bookstore_oauth_api/src/utils/errors"
	"github.com/shchaslyvyi/bookstore_users-api/utils/crypto_utils"
)

const (
	expirationTime             = 24
	grantTypePasswort          = "passwort"
	grantTypeClientCredentials = "client_credentials"
)

// AccessTokenRequest struct defines an access token
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant_type
	Username string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate func validates the Access Token Request
func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePasswort:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("Invalid grant type parametr.")
	}
	// TADO: validate parameters for each grant_type.
	return nil
}

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
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired function returns the validity state of the access token
func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

// Generate function returns generatedaccess token
func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
