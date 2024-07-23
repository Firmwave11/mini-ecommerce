package models

import "time"

type OrderProduct struct {
	ID          int        `json:"id" db:"id"`
	OrderId     int        `json:"order_id" db:"order_id"`
	ProductId   int        `json:"product_id" db:"product_id"`
	WarehouseId int        `json:"warehouse_id" db:"warehouse_id"`
	Quantity    int        `json:"total_price" db:"total_price"`
	Weight      float64    `json:"weight" db:"weight"`
	Price       float64    `json:"Price" db:"Price"`
	CreatedAT   *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAT   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
}
