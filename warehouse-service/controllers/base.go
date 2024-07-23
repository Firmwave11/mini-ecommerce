package controllers

import (
	warehouseService "warehouse-service/services"

	"github.com/labstack/echo/v4"
)

type controller struct {
	warehouseService warehouseService.WarehouseService
}

type Controller interface {
	GetStockWarehouseHandler(e echo.Context) error
	UpdateStockWarehouseHandler(e echo.Context) error
	SyncStockWarehouseHandler(e echo.Context) error
	GetWarehouseHandler(e echo.Context) error
	InsertStockWarehouseHandler(e echo.Context) error
	InsertWarehouseHandler(e echo.Context) error
}

func NewController(warehouse warehouseService.WarehouseService) Controller {
	return &controller{
		warehouseService: warehouse,
	}
}
