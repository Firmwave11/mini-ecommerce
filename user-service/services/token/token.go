package token

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
	model "user-service/models/token"
	utils "user-service/utils/encrypt"
	utilsEncrypt "user-service/utils/encrypt"
	utilsLog "user-service/utils/log"
	"user-service/utils/response"

	"github.com/spf13/viper"
)

var (
	// ErrLoginFail :nodoc:
	ErrLoginFail = errors.New("login fail")
	// ErrUserNotFound :nodoc:
	ErrUserNotFound = errors.New("user not found")
	// ErrAccountDisabled :nodoc:
	ErrTokenNotValid = errors.New("token not valid")
	ErrGenerateToken = errors.New("invalid generate token")
)

func (t *tokenService) GenerateToken(ctx context.Context, req model.Login) (*model.UserToken, error) {

	res, err := t.tokenRepo.GetUserAccountByEmailorPhone(req.Input)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid email or phone or password.\r\n")
		}

		return nil, err
	}

	isPasswordVerified := utils.Verify(req.Password, res.Salt, res.Password)

	if !isPasswordVerified {
		return nil, ErrLoginFail
	}

	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 60)
	key := viper.GetString("application.secret")

	payload := model.Payload{
		Type:      "access_token",
		Scope:     "public",
		UserID:    res.ID.String(),
		CreatedAt: dt,
		ExpiredAt: expirtyTime,
	}

	token, err := utilsEncrypt.Encrypt(payload, key)
	if err != nil {
		utilsLog.LogError(err)
		return nil, ErrGenerateToken
	}

	accessToken := base64.StdEncoding.EncodeToString(token)
	modelToken := model.UserToken{}
	modelToken.AuthToken = accessToken
	modelToken.CreatedAt = dt
	modelToken.UserAccountID = res.ID
	modelToken.ExpiredAt = expirtyTime

	resToken, err := t.tokenRepo.InsertUserToken(&modelToken)
	if err != nil {
		return nil, err
	}
	modelToken.TokenID = resToken.TokenID
	return &modelToken, nil

}

func (t *tokenService) ValidateToken(ctx context.Context, auth_token string) response.Output {
	res, err := t.tokenRepo.GetUserAuthFromTokenRequest(auth_token)
	dt := time.Now()
	if err != nil {

		if err == sql.ErrNoRows {
			return response.Errors(ctx, http.StatusUnauthorized, "000003", "Failed Check Token", "Token Not Valid", ErrTokenNotValid)
		}

		return response.Errors(ctx, http.StatusUnauthorized, "000003", "Failed Check Token", "Check Token Error", err)
	}

	if res.ExpiredAt.Before(dt) {
		return response.Errors(ctx, http.StatusUnauthorized, "000003", "Failed Check Token", "Token Expired", errors.New("The token is expired.\r\n"))
	}

	return response.Success(ctx, http.StatusOK, "000003", "Token Valid", res)
}

func (t *tokenService) LoginAction(ctx context.Context, req model.Login) response.Output {
	res, err := t.GenerateToken(ctx, req)
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Login", "Failed Login", err)
	}

	return response.Success(ctx, http.StatusCreated, "000002", "Success Login", res)
}

func (w *tokenService) User(ctx context.Context, userId string) response.Output {
	fmt.Println(userId)
	res, err := w.tokenRepo.GetProfileByUserId(userId)
	if err != nil && err != sql.ErrNoRows {
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Get Customer", "Failed Get Customer", err)
	}

	return response.Success(ctx, http.StatusOK, "000004", "Success Get Customer", res)
}
