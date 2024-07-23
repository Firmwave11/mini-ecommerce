package services

import (
	"context"
	"database/sql"
	"net/http"
	models "shop-service/models/entity"
	"shop-service/utils/response"
)

func (w *shopService) GetShop(ctx context.Context) response.Output {

	res, err := w.shopRepo.FindAllShop()
	if err != nil && err != sql.ErrNoRows {
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Get Shop", "Failed Get Shop", err)
	}

	return response.Success(ctx, http.StatusOK, "000001", "Success Get Shop", res)
}

func (w *shopService) InsertorUpdateShop(ctx context.Context, req models.Shop) response.Output {
	tx := w.db.MustBegin()

	err := w.shopRepo.InsertOrUpdateShop(tx, req)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Insert or Update Shop", "Failed Insert or Update Shop", err)
	}

	err = tx.Commit()

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Insert or Update Shop", "Failed Insert or Update Shop", err)
	}

	return response.Success(ctx, http.StatusCreated, "000002", "Success Insert or Update Shop", nil)
}
