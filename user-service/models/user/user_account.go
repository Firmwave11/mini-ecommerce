package user

import (
	"time"

	"github.com/google/uuid"
)

// UserAccount --
type UserAccount struct {
	ID        uuid.UUID  `db:"id"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	Salt      string     `db:"salt"`
	Name      string     `db:"name"`
	Phone     string     `db:"phone"`
	CreatedAT *time.Time `db:"created_at"`
	UpdatedAT *time.Time `db:"updated_at"`
	IsDeleted bool       `db:"is_deleted"`
}
