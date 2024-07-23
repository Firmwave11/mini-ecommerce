package repositories

import (
	"warehouse-service/middleware"
	models "warehouse-service/models/entity"

	"github.com/jmoiron/sqlx"
)

var timezone = middleware.Timezone()

type warehouseRepo struct {
	db *sqlx.DB
}

type WarehouseRepo interface {
	InsertOrUpdateStockWarehouse(tx *sqlx.Tx, stockReq models.Stock) error
	FindWarehouseEnable() ([]*models.Warehouse, error)
	FindStockProductWarehouseEnable(productIds []int) ([]*models.Stock, error)
	UpdateDecreaseStockWarehouse(tx *sqlx.Tx, stockReq models.Stock) error
	UpdateIncreaseStockWarehouse(tx *sqlx.Tx, stockReq models.Stock) error
	FindStocks(productId int, warehouseId int) ([]*models.Stock, error)
	InsertOrUpdateWarehouse(tx *sqlx.Tx, warehouseReq models.Warehouse) error
}

func NewWarehouseRepository(db *sqlx.DB) WarehouseRepo {
	return &warehouseRepo{
		db: db,
	}
}
