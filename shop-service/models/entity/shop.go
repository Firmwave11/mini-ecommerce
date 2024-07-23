package models

import "time"

type Shop struct {
	ID         int        `json:"id" db:"id"`
	CustomerId int        `json:"customer_id" db:"customer_id"`
	Name       string     `json:"name" db:"name"`
	Address    *string    `json:"address,omitempty" db:"address"`
	Phone      *string    `json:"phone,omitempty" db:"phone"`
	CreatedAT  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAT  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	IsDeleted  bool       `json:"is_deleted" db:"is_deleted"`
}
