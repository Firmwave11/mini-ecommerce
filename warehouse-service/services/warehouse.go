package services

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	models "warehouse-service/models/entity"
	"warehouse-service/models/request"
	"warehouse-service/utils/response"
)

func (w *warehouseService) GetStockWarehouse(ctx context.Context, req request.Stock) response.Output {

	res, err := w.warehouseRepo.FindStockProductWarehouseEnable(req.ProductId)
	if err != nil && err != sql.ErrNoRows {
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Get Stock Warehouse", "Failed Get Stock Warehouse", err)
	}

	return response.Success(ctx, http.StatusOK, "000001", "Success Get Stock Warehouse", res)
}

func (w *warehouseService) UpdateStockWarehouse(ctx context.Context, req request.UpdateStockRequest) response.Output {
	tx := w.db.MustBegin()
	stock := models.Stock{
		Quantity:    req.Quantity,
		WarehouseId: req.WarehouseId,
		IsDeleted:   false,
		ProductId:   req.ProductId,
	}
	if req.Type == "decrease" {

		res, err := w.warehouseRepo.FindStocks(stock.ProductId, stock.WarehouseId)
		if err != nil {
			tx.Rollback()
			return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Update Stock Warehouse", "Failed Update Stock Warehouse", err)
		}

		if res[len(res)-1].Quantity < req.Quantity {
			tx.Rollback()
			return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Update Stock Warehouse", "Stok Less Than Request", errors.New("Stok Less Than Request"))
		}

		err = w.warehouseRepo.UpdateDecreaseStockWarehouse(tx, stock)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Update Stock Warehouse", "Failed Update Stock Warehouse", err)
		}
	} else {
		err := w.warehouseRepo.UpdateIncreaseStockWarehouse(tx, stock)
		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Update Stock Warehouse", "Failed Update Stock Warehouse", err)
		}
	}
	err := tx.Commit()

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Update Stock Warehouse", "Failed Update Stock Warehouse", err)
	}

	return response.Success(ctx, http.StatusCreated, "000002", "Success Update Stock Warehouse", nil)
}

func (w *warehouseService) SyncStockWarehouse(ctx context.Context, req request.UpdateSyncStockWarehouseRequest) response.Output {
	tx := w.db.MustBegin()
	stockFrom := models.Stock{
		Quantity:    req.Quantity,
		WarehouseId: req.FromWarehouseId,
		IsDeleted:   false,
		ProductId:   req.ProductId,
	}

	stockTo := models.Stock{
		Quantity:    req.Quantity,
		WarehouseId: req.ToWarehouseId,
		IsDeleted:   false,
		ProductId:   req.ProductId,
	}

	res, err := w.warehouseRepo.FindStocks(stockFrom.ProductId, stockFrom.WarehouseId)
	if err != nil {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000003", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	if res[len(res)-1].Quantity < req.Quantity {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000003", "Failed Insert or Update Stock Warehouse", "Stok Less Than Request", errors.New("Stok Less Than Request"))
	}

	err = w.warehouseRepo.UpdateDecreaseStockWarehouse(tx, stockFrom)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000003", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	err = w.warehouseRepo.UpdateIncreaseStockWarehouse(tx, stockTo)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000003", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	err = tx.Commit()

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000003", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	return response.Success(ctx, http.StatusCreated, "000003", "Success Insert or Update Stock Warehouse", nil)
}

func (w *warehouseService) InsertStockWarehouse(ctx context.Context, req models.Stock) response.Output {
	tx := w.db.MustBegin()
	stock := models.Stock{
		Quantity:    req.Quantity,
		WarehouseId: req.WarehouseId,
		IsDeleted:   false,
		ProductId:   req.ProductId,
	}

	err := w.warehouseRepo.InsertOrUpdateStockWarehouse(tx, stock)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	err = tx.Commit()

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	return response.Success(ctx, http.StatusCreated, "000004", "Success Insert or Update Stock Warehouse", nil)
}

func (w *warehouseService) InsertWarehouse(ctx context.Context, req models.Warehouse) response.Output {
	tx := w.db.MustBegin()

	err := w.warehouseRepo.InsertOrUpdateWarehouse(tx, req)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000005", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	err = tx.Commit()

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000005", "Failed Insert or Update Stock Warehouse", "Failed Insert or Update Stock Warehouse", err)
	}

	return response.Success(ctx, http.StatusCreated, "000005", "Success Insert or Update Stock Warehouse", nil)
}

func (w *warehouseService) GetWarehouse(ctx context.Context) response.Output {
	res, err := w.warehouseRepo.FindWarehouseEnable()
	if err != nil && err != sql.ErrNoRows {
		return response.Errors(ctx, http.StatusBadRequest, "000006", "Failed Get Warehouse Enable", "Failed Get Warehouse Enable", err)
	}

	return response.Success(ctx, http.StatusOK, "000006", "Success Get Warehouse", res)
}
