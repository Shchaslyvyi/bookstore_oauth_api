package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shchaslyvyi/bookstore_oauth_api/src/domain/access_token"

	"github.com/shchaslyvyi/bookstore_oauth_api/src/utils/errors"

	res "github.com/shchaslyvyi/bookstore_oauth_api/src/services/access_token"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

// AccessTokenHandler structure
type accessTokenHandler struct {
	service res.Service
}

// NewAccessTokenHandler is a func returning the handler
func NewAccessTokenHandler(service res.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := handler.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body.")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
