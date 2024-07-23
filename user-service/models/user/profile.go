package user

import (
	"time"

	"github.com/google/uuid"
)

// Profiles --
type Profiles struct {
	ID            int64      `db:"id"`
	FirstName     string     `db:"first_name"`
	LastName      string     `db:"last_name"`
	Gender        string     `db:"gender"`
	BirthDate     time.Time  `db:"birth_date"`
	PhoneNumber   string     `db:"phone_number"`
	Nickname      string     `db:"nickname"`
	Description   *string    `db:"description"`
	Photo         *string    `db:"photo"`
	UpdatedAT     *time.Time `db:"updated_at"`
	CreatedAT     *time.Time `db:"created_at"`
	UserAccountID uuid.UUID  `db:"user_account_id"`
	IsDeleted     bool       `db:"is_deleted"`
	IsPremium     bool       `db:"is_premium"`
	IsVerified    bool       `db:"is_verified"`
}
