package token

import (
	model "user-service/models/token"
	"user-service/models/user"
	modelUser "user-service/models/user"

	"github.com/jmoiron/sqlx"
)

type tokenRepo struct {
	db *sqlx.DB
}

type TokenRepo interface {
	GetUserAuthFromTokenRequest(authToken string) (userToken model.UserToken, err error)
	GetUserAccountByEmailorPhone(input string) (user modelUser.UserAccount, err error)
	InsertUserToken(userToken *model.UserToken) (*model.UserToken, error)
	GetProfileByUserId(userId string) (profile user.Profiles, err error)
}

func NewTokenRepository(db *sqlx.DB) TokenRepo {
	return &tokenRepo{
		db: db,
	}
}
