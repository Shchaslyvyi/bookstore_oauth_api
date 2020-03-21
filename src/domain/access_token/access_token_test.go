package access_token

import (
	"testing"
	"time"

	"github.com/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration Time should be 24 hours.")
}
func TestGetNowAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "New access token should not be nil.")
	assert.EqualValues(t, "", at.AccessToken, "New access token should not have defined access token id.")
	assert.True(t, at.UserID == 0, "New access token should not have an associated user id.")
}
func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default.")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Empty access token expiring 3 hours from now should not be expired.")
}
