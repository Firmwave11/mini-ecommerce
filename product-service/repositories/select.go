package repositories

import (
	models "product-service/models/entity"
)

func (w *productRepo) FindAllProduct() ([]*models.Product, error) {
	var (
		shop = make([]*models.Product, 0)
		err  error
	)
	query := `SELECT id, "name", price, weight, enabled, updated_at, created_at, is_deleted FROM public.product where is_deleted = false`

	err = w.db.Select(&shop, query)
	return shop, err
}
