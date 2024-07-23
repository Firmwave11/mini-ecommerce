package model

import (
	"time"

	"github.com/google/uuid"
)

type UserToken struct {
	TokenID       uuid.UUID `db:"token_id"`
	UserAccountID uuid.UUID `db:"user_account_id"`
	AuthToken     string    `db:"auth_token"`
	CreatedAt     time.Time `db:"generated_at"`
	ExpiredAt     time.Time `db:"expires_at"`
}
