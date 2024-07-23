package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	models "warehouse-service/models/entity"

	"github.com/jmoiron/sqlx"
)

func (w *warehouseRepo) InsertOrUpdateStockWarehouse(tx *sqlx.Tx, stockReq models.Stock) error {
	var (
		stock models.Stock
		args  []interface{}
	)
	query := `SELECT id, warehouse_id, product_id, quantity, updated_at, created_at, is_deleted FROM public.stock where is_deleted = false and id = $1`
	err := tx.Get(&stock, query, stockReq.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO public.stock (warehouse_id, product_id, quantity, updated_at, created_at, is_deleted) VALUES($1, $2, $3, now(), now(), $4)`
			args = []interface{}{
				stockReq.WarehouseId,
				stockReq.ProductId,
				stockReq.Quantity,
				stockReq.IsDeleted,
			}
		} else {
			return err
		}
	} else {
		query = `UPDATE public.stock SET warehouse_id=$1, product_id=$2, quantity=$3, updated_at=now(), created_at=now(), is_deleted=$5 WHERE id=$4
`
		args = []interface{}{
			stockReq.WarehouseId,
			stockReq.ProductId,
			stockReq.Quantity,
			stockReq.IsDeleted,
			stockReq.ID,
		}
	}

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	return err
}

func (w *warehouseRepo) UpdateDecreaseStockWarehouse(tx *sqlx.Tx, stockReq models.Stock) error {

	queryString := "lock table stock in access exclusive mode"

	_, err := tx.Exec(queryString)

	query := `update public.stock SET quantity= quantity - $1, updated_at=now() where product_id = $2 and warehouse_id = $3 and updated_at < now()`

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		stockReq.Quantity,
		stockReq.ProductId,
		stockReq.WarehouseId,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if rowsAffected != 1 {
		err = fmt.Errorf("failed update data event_registration. rows affected %d", rowsAffected)
		return err
	}

	return err
}

func (w *warehouseRepo) UpdateIncreaseStockWarehouse(tx *sqlx.Tx, stockReq models.Stock) error {

	queryString := "lock table stock in access exclusive mode"

	_, err := tx.Exec(queryString)

	if err != nil {
		return err
	}
	query := `update public.stock SET quantity= quantity + $1, updated_at=now() where product_id = $2 and warehouse_id = $3 and updated_at < now()`

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		stockReq.Quantity,
		stockReq.ProductId,
		stockReq.WarehouseId,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if rowsAffected != 1 {
		err = fmt.Errorf("failed update data event_registration. rows affected %d", rowsAffected)
		return err
	}

	return err
}

func (w *warehouseRepo) InsertOrUpdateWarehouse(tx *sqlx.Tx, warehouseReq models.Warehouse) error {
	var (
		warehouse models.Warehouse
		args      []interface{}
	)
	query := `SELECT id, "name", shop_id, enabled, area, address, updated_at, created_at, is_deleted FROM public.warehouse where is_deleted = false and id = $1`
	err := tx.Get(&warehouse, query, warehouseReq.ID)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO public.warehouse ("name", enabled, area, address, updated_at, created_at, is_deleted, shop_id) VALUES($1, $2, $3, $4, now(), now(), $5, $6)`
			args = []interface{}{
				warehouseReq.Name,
				warehouseReq.Enabled,
				warehouseReq.Area,
				warehouseReq.Address,
				warehouseReq.IsDeleted,
				warehouseReq.ShopId,
			}
		} else {
			return err
		}
	} else {
		query = `UPDATE public.warehouse SET "name"=$1, enabled=$2, area=$3, address=$4, updated_at=now(), is_deleted=$5, shop_id=$6 WHERE id=$7`
		args = []interface{}{
			warehouseReq.Name,
			warehouseReq.Enabled,
			warehouseReq.Area,
			warehouseReq.Address,
			warehouseReq.IsDeleted,
			warehouseReq.ShopId,
			warehouseReq.ID,
		}
	}

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	return err
}
