package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/domain/access_token"
)

// AccessTokenHandler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

// AccessTokenHandler structure
type accessTokenHandler struct {
	service access_token.Service
}

// NewHandler is a func returning the handler
func NewHandler(service access_token.Service) AccessTokenHandler {
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
	c.JSON(http.StatusNotImplemented, accessToken)
}
