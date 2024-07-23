package models

import "time"

type Order struct {
	ID            int        `json:"id" db:"id"`
	CustomerId    int        `json:"customer_id" db:"customer_id"`
	TotalQuantity int        `json:"total_quantity" db:"total_quantity"`
	TotalWeight   float64    `json:"total_weight" db:"total_weight"`
	TotalPrice    float64    `json:"total_price" db:"total_price"`
	IsPaymented   bool       `json:"is_paymented" db:"is_paymented"`
	CreatedAT     *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAT     *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	IsDeleted     bool       `json:"is_deleted" db:"is_deleted"`
}
