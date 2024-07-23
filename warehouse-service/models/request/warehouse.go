package request

import "time"

type Warehouse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Enabled   bool       `json:"enabled"`
	Address   *string    `json:"address"`
	Area      *string    `json:"area"`
	CreatedAT *time.Time `json:"created_at"`
	UpdatedAT *time.Time `json:"updated_at"`
	IsDeleted bool       `json:"is_deleted"`
	ShopId    int        `json:"shop_id"`
}
