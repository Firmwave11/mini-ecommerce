package services

import (
	"context"
	models "warehouse-service/models/entity"
	"warehouse-service/models/request"
	warehouseRepo "warehouse-service/repositories"
	"warehouse-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type warehouseService struct {
	warehouseRepo warehouseRepo.WarehouseRepo
	db            *sqlx.DB
}

type WarehouseService interface {
	GetStockWarehouse(ctx context.Context, req request.Stock) response.Output
	UpdateStockWarehouse(ctx context.Context, req request.UpdateStockRequest) response.Output
	SyncStockWarehouse(ctx context.Context, req request.UpdateSyncStockWarehouseRequest) response.Output
	InsertWarehouse(ctx context.Context, req models.Warehouse) response.Output
	InsertStockWarehouse(ctx context.Context, req models.Stock) response.Output
	GetWarehouse(ctx context.Context) response.Output
}

func NewWarehouseService(e warehouseRepo.WarehouseRepo, db *sqlx.DB) WarehouseService {
	return &warehouseService{
		warehouseRepo: e,
		db:            db,
	}
}
