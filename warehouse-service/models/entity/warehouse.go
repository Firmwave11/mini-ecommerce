package models

import "time"

type Warehouse struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Enabled   bool       `json:"enabled" db:"enabled"`
	Address   *string    `json:"address,omitempty" db:"address"`
	Area      *string    `json:"area,omitempty" db:"area"`
	CreatedAT *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAT *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	IsDeleted bool       `json:"is_deleted" db:"is_deleted"`
	ShopId    int        `json:"shop_id" db:"shop_id"`
}
