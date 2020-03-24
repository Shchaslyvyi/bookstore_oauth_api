package app

import (
	"github.com/gin-gonic/gin"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/clients/cassandra"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/domain/access_token"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/repository/db"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/repository/http"
)

var (
	router = gin.Default()
)

// StartApplication func starts the new application
func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
