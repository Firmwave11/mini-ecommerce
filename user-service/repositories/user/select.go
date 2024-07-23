package user

import (
	"user-service/models/user"
)

func (t *userRepo) GetProfileByUserId(userId string) (profile user.Profiles, err error) {

	query := "SELECT id, first_name, last_name, gender, birth_date, phone_number, nickname, description, photo, updated_at, created_at, is_deleted, user_account_id, is_verified FROM public.customer where user_account_id = $1 "

	err = t.db.QueryRowx(query, userId).StructScan(&profile)

	return profile, err
}
