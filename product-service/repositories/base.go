package repositories

import (
	"product-service/middleware"
	models "product-service/models/entity"

	"github.com/jmoiron/sqlx"
)

var timezone = middleware.Timezone()

type productRepo struct {
	db *sqlx.DB
}

type ProductRepo interface {
	InsertOrUpdateProduct(tx *sqlx.Tx, productReq models.Product) error
	FindAllProduct() ([]*models.Product, error)
}

func NewProductRepository(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}
