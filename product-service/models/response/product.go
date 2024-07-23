package models

import (
	models "product-service/models/entity"
	"time"
)

type ProductResponse struct {
	ID        int             `json:"id" db:"id"`
	Name      string          `json:"name" db:"name"`
	Price     float64         `json:"price" db:"price"`
	Weight    float64         `json:"weight,omitempty" db:"weight"`
	Enabled   bool            `json:"enabled" db:"enabled"`
	CreatedAT *time.Time      `json:"created_at,omitempty" db:"created_at"`
	UpdatedAT *time.Time      `json:"updated_at,omitempty" db:"updated_at"`
	IsDeleted bool            `json:"is_deleted" db:"is_deleted"`
	Stock     []*models.Stock `json:"warehouse"`
}
