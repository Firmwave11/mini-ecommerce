package models

import (
	"time"
)

type Stock struct {
	ID          int        `json:"id" db:"id"`
	WarehouseId int        `json:"warehouse_id" db:"warehouse_id"`
	ProductId   int        `json:"product_id" db:"product_id"`
	Quantity    int        `json:"quantity" db:"quantity"`
	CreatedAT   *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAT   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	Type        string     `json:"type"`
}
