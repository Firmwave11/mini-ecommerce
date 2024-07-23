package repositories

import (
	"database/sql"
	"errors"
	models "shop-service/models/entity"

	"github.com/jmoiron/sqlx"
)

func (w *shopRepo) InsertOrUpdateShop(tx *sqlx.Tx, shopReq models.Shop) error {
	var (
		shop models.Shop
		args []interface{}
	)
	query := `SELECT id, customer_id, address, "name", phone, updated_at, created_at, is_deleted FROM public.shop where is_deleted = false and id = $1`
	err := tx.Get(&shop, query, shopReq.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query = `INSERT INTO public.shop (customer_id, address, "name", phone, updated_at, created_at, is_deleted) VALUES($1, $2, $3, $4, now(), now(), $5);`
			args = []interface{}{
				shopReq.CustomerId,
				shopReq.Address,
				shopReq.Name,
				shopReq.Phone,
				shopReq.IsDeleted,
			}
		} else {
			return err
		}
	} else {
		query = `UPDATE public.shop SET customer_id=$1, address=$2, "name"=$3, phone=$4, updated_at=now(), is_deleted=$5 WHERE id=$6`
		args = []interface{}{
			shopReq.CustomerId,
			shopReq.Address,
			shopReq.Name,
			shopReq.Phone,
			shopReq.IsDeleted,
			shopReq.ID,
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
