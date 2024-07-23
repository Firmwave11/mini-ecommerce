package repositories

import (
	"database/sql"
	"errors"
	models "product-service/models/entity"

	"github.com/jmoiron/sqlx"
)

func (w *productRepo) InsertOrUpdateProduct(tx *sqlx.Tx, productReq models.Product) error {
	var (
		product models.Product
		args    []interface{}
	)
	query := `SELECT id, "name", price, weight, enabled, updated_at, created_at, is_deleted FROM public.product where is_deleted = false and id = $1`
	err := tx.Get(&product, query, productReq.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO public.product ("name", price, weight, enabled, updated_at, created_at, is_deleted) VALUES($1, $2, $3, $4, now(), now(), $5)`
			args = []interface{}{
				productReq.Name,
				productReq.Price,
				productReq.Weight,
				productReq.Enabled,
				productReq.IsDeleted,
			}
		} else {
			return err
		}
	} else {
		query = `UPDATE public.product SET "name"=$1, price=$2, weight=$3, enabled=$4, updated_at=now(), is_deleted=$5 WHERE id=$6`
		args = []interface{}{
			productReq.Name,
			productReq.Price,
			productReq.Weight,
			productReq.Enabled,
			productReq.IsDeleted,
			productReq.ID,
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
