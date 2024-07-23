package repositories

import (
	models "order-service/models/entity"
)

func (w *orderRepo) FindAllOrder() ([]*models.Order, error) {
	var (
		shop = make([]*models.Order, 0)
		err  error
	)
	query := `SELECT id, customer_id, address, "name", phone, updated_at, created_at, is_deleted FROM public.shop where is_deleted = false and is_deleted = false`

	err = w.db.Select(&shop, query)
	return shop, err
}
