package token

import (
	modelToken "user-service/models/token"
	"user-service/models/user"
	modelUser "user-service/models/user"
)

func (t *tokenRepo) GetUserAuthFromTokenRequest(authToken string) (userToken modelToken.UserToken, err error) {

	query := "SELECT token_id, user_account_id, auth_token, generated_at, expires_at FROM public.authentication_tokens WHERE auth_token = $1"

	err = t.db.QueryRowx(query, authToken).StructScan(&userToken)

	return userToken, err
}

func (t *tokenRepo) GetUserAccountByEmailorPhone(input string) (user modelUser.UserAccount, err error) {
	query := `SELECT id, email, "password", salt, "name", phone, created_at, updated_at, is_deleted FROM public.user_account where email = $1 or phone = $1`

	err = t.db.QueryRowx(query, input).StructScan(&user)

	return user, err
}

func (t *tokenRepo) GetProfileByUserId(userId string) (profile user.Profiles, err error) {

	query := "SELECT id, first_name, last_name, gender, birth_date, phone_number, nickname, description, photo, updated_at, created_at, is_deleted, user_account_id, is_verified FROM public.customer where user_account_id = $1 "

	err = t.db.QueryRowx(query, userId).StructScan(&profile)

	return profile, err
}
