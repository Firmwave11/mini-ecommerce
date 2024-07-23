package repositories

import (
	models "shop-service/models/entity"
)

func (w *shopRepo) FindAllShop() ([]*models.Shop, error) {
	var (
		shop = make([]*models.Shop, 0)
		err  error
	)
	query := `SELECT id, customer_id, address, "name", phone, updated_at, created_at, is_deleted FROM public.shop where is_deleted = false and is_deleted = false`

	err = w.db.Select(&shop, query)
	return shop, err
}
