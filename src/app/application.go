package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/domain/access_token"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/repository/db"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/repository/http"
)

var (
	router = gin.Default()
)

// StartApplication func starts the new application
func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8080")
}
