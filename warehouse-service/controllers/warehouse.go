package controllers

import (
	models "warehouse-service/models/entity"
	"warehouse-service/models/request"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (c *controller) GetStockWarehouseHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqStock = &request.Stock{}
	)

	// creates query params binder that stops binding at first error
	err := echo.QueryParamsBinder(e).
		Ints("productIds", &reqStock.ProductId).
		BindError()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "GetStockWarehouseHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.warehouseService.GetStockWarehouse(ctx, *reqStock).Write(e)
}

func (c *controller) UpdateStockWarehouseHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqStock = &request.UpdateStockRequest{}
	)

	if err := e.Bind(&reqStock); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "UpdateStockWarehouseHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.warehouseService.UpdateStockWarehouse(ctx, *reqStock).Write(e)
}

func (c *controller) SyncStockWarehouseHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqStock = &request.UpdateSyncStockWarehouseRequest{}
	)

	if err := e.Bind(&reqStock); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "UpdateStockWarehouseHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.warehouseService.SyncStockWarehouse(ctx, *reqStock).Write(e)
}

func (c *controller) GetWarehouseHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
	)

	return c.warehouseService.GetWarehouse(ctx).Write(e)
}

func (c *controller) InsertStockWarehouseHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqStock = &models.Stock{}
	)

	if err := e.Bind(&reqStock); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "InsertStockWarehouseHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.warehouseService.InsertStockWarehouse(ctx, *reqStock).Write(e)
}

func (c *controller) InsertWarehouseHandler(e echo.Context) error {
	var (
		req          = e.Request()
		ctx          = req.Context()
		reqWarehouse = &models.Warehouse{}
	)

	if err := e.Bind(&reqWarehouse); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "InsertWarehouseHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.warehouseService.InsertWarehouse(ctx, *reqWarehouse).Write(e)
}
