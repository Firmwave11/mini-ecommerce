package user

import (
	"context"
	model "user-service/models/user"
	tokenRepo "user-service/repositories/token"
	"user-service/repositories/user"
	"user-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type userService struct {
	userRepo  user.UserRepo
	tokenRepo tokenRepo.TokenRepo
	db        *sqlx.DB
}

type UserService interface {
	RegisterAction(ctx context.Context, req model.Register) response.Output
}

func NewUserService(tokenRepo tokenRepo.TokenRepo, user user.UserRepo, db *sqlx.DB) UserService {
	return &userService{
		tokenRepo: tokenRepo,
		userRepo:  user,
		db:        db,
	}
}
