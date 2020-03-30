package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/http"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/repository/db"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/repository/rest"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/services/access_token"
)

var (
	router = gin.Default()
)

// StartApplication func starts the new application
func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRestUserRepository(), db.NewRepository()))
	router.GET("oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
