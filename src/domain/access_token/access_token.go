package access_token

import (
	"time"
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
