package repositories

import (
	models "warehouse-service/models/entity"

	"github.com/jmoiron/sqlx"
)

func (w *warehouseRepo) FindWarehouseEnable() ([]*models.Warehouse, error) {
	var (
		stocks = make([]*models.Warehouse, 0)
		err    error
	)
	query := `SELECT id, "name", enabled, area, address, updated_at, created_at, is_deleted, shop_id FROM public.warehouse where enabled = true and is_deleted = false`

	err = w.db.Select(&stocks, query)
	return stocks, err
}

func (w *warehouseRepo) FindStockProductWarehouseEnable(productIds []int) ([]*models.Stock, error) {
	var (
		stocks = make([]*models.Stock, 0)
		err    error
	)

	q, args, err := sqlx.In(`select s.id, s.product_id , s.quantity, s.warehouse_id, s.updated_at, s.created_at
	from public.stock s 
	inner join warehouse w ON w.id = s.warehouse_id and w.enabled = true and w.is_deleted = false
	where s.is_deleted = false and s.product_id in (?)`, productIds)
	if err != nil {
		return nil, err
	}
	err = w.db.Select(&stocks, sqlx.Rebind(sqlx.DOLLAR, q), args...)
	return stocks, err
}

func (w *warehouseRepo) FindStocks(productId int, warehouseId int) ([]*models.Stock, error) {
	var (
		stocks = make([]*models.Stock, 0)
		err    error
	)
	query := `SELECT id, warehouse_id, product_id, quantity, updated_at, created_at, is_deleted FROM public.stock where is_deleted = false and warehouse_id = $1 and product_id = $2`

	err = w.db.Select(&stocks, query, productId, warehouseId)
	return stocks, err
}
