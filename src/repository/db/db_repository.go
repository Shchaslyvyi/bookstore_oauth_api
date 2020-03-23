package db

import (
	"github.com/shchaslyvyi/bookstore_oauth_api/src/clients/cassandra"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/domain/access_token"
	"github.com/shchaslyvyi/bookstore_oauth_api/src/utils/errors"
)

// DbRepository interface
type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

// NewRepository function to access an dbRepository struct
func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(ID string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("DB connection not implemented yet.")
}
