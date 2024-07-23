package token

import (
	"time"
	model "user-service/models/token"
)

func (t *tokenRepo) InsertUserToken(userToken *model.UserToken) (*model.UserToken, error) {
	now := time.Now()
	query := `
	INSERT INTO public.authentication_tokens (user_account_id, auth_token, generated_at, expires_at) VALUES( $1, $2, $3, $4)
	returning token_Id
	`

	err := t.db.DB.QueryRow(
		query,
		userToken.UserAccountID,
		userToken.AuthToken,
		now,
		userToken.ExpiredAt,
	).Scan(&userToken.TokenID)

	return userToken, err
}
