package repositories

import (
	"database/sql"
	"errors"
	models "order-service/models/entity"

	"github.com/jmoiron/sqlx"
)

func (w *orderRepo) InsertOrUpdateOrder(tx *sqlx.Tx, orderReq models.Order) (int, error) {
	var (
		order models.Order
		args  []interface{}
		id    int
	)
	query := `SELECT id, customer_id, total_quantity, total_weight, total_price, is_paymented, updated_at, created_at, is_deleted FROM public."order" where is_deleted = false and id = $1`
	err := tx.Get(&order, query, orderReq.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO public."order" (customer_id, total_quantity, total_weight, total_price, is_paymented, updated_at, created_at, is_deleted) VALUES($1, $2, $3, $4, $5, now(), now(), $6) returning id`
			args = []interface{}{
				orderReq.CustomerId,
				orderReq.TotalQuantity,
				orderReq.TotalWeight,
				orderReq.TotalPrice,
				orderReq.IsPaymented,
				orderReq.IsDeleted,
			}
		} else {
			return 0, err
		}
	} else {
		query = `UPDATE public."order" SET customer_id=$1, total_quantity=$2, total_weight=$3, total_price=$4, is_paymented=$5, updated_at=now(), is_deleted=$7 WHERE id=$7`
		args = []interface{}{
			orderReq.CustomerId,
			orderReq.TotalQuantity,
			orderReq.TotalWeight,
			orderReq.TotalPrice,
			orderReq.IsPaymented,
			orderReq.IsDeleted,
			orderReq.ID,
		}
	}

	stmt, err := tx.Preparex(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRowx(args...).Scan(id)
	if err != nil {
		return 0, err
	}
	return 0, err
}

func (w *orderRepo) InsertOrUpdateOrderProduct(tx *sqlx.Tx, orderProductReq models.OrderProduct) (int, error) {
	var (
		orderProduct models.OrderProduct
		args         []interface{}
		id           int
	)
	query := `SELECT id, product_id, order_id, warehouse_id, quantity, weight, price, updated_at, created_at, is_deleted FROM public.order_product where is_deleted = false and id = $1`
	err := tx.Get(&orderProduct, query, orderProductReq.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO public.order_product (product_id, order_id, warehouse_id, quantity, weight, price, updated_at, created_at, is_deleted) VALUES($1, $2, $3, $4, $5, $6, now(), now(), $7) returning id`
			args = []interface{}{
				orderProductReq.ProductId,
				orderProductReq.OrderId,
				orderProduct.WarehouseId,
				orderProductReq.Quantity,
				orderProductReq.Weight,
				orderProductReq.Price,
				orderProductReq.UpdatedAT,
				orderProductReq.CreatedAT,
				orderProductReq.IsDeleted,
			}
		} else {
			return 0, err
		}
	} else {
		query = `UPDATE public.order_product SET product_id=$1, order_id=$2, warehouse_id=$3, quantity=$4, weight=$5, price=$6, updated_at=now(), is_deleted=$7 WHERE id=$8`
		args = []interface{}{
			orderProductReq.ProductId,
			orderProductReq.OrderId,
			orderProduct.WarehouseId,
			orderProductReq.Quantity,
			orderProductReq.Weight,
			orderProductReq.Price,
			orderProductReq.UpdatedAT,
			orderProductReq.CreatedAT,
			orderProductReq.IsDeleted,
		}
	}

	stmt, err := tx.Preparex(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRowx(args...).Scan(id)
	if err != nil {
		return 0, err
	}
	return id, err
}
