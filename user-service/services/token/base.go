package token

import (
	"context"
	model "user-service/models/token"
	tokenRepo "user-service/repositories/token"
	"user-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type tokenService struct {
	tokenRepo tokenRepo.TokenRepo
	db        *sqlx.DB
}

type TokenService interface {
	ValidateToken(ctx context.Context, auth_token string) response.Output
	GenerateToken(ctx context.Context, req model.Login) (*model.UserToken, error)
	LoginAction(ctx context.Context, req model.Login) response.Output
	User(ctx context.Context, userId string) response.Output
}

func NewTokenService(e tokenRepo.TokenRepo, db *sqlx.DB) TokenService {
	return &tokenService{
		tokenRepo: e,
		db:        db,
	}
}
