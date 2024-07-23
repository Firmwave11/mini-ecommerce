package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/logger"
	models "order-service/models/entity"
	req "order-service/models/request"
	"order-service/utils/request"
	"order-service/utils/response"

	"github.com/spf13/viper"
)

func (w *orderService) CreateOrder(ctx context.Context, reqs *[]req.ProductRequest, userAccountId string, token string) response.Output {
	var (
		totalPrice    float64
		TotalWeight   float64
		TotalQuantity int
		order         = &models.Order{}
		customer      = &models.Customer{}
		warehouseHost = viper.GetString("client.warehouse.host")
		urlWarehouse  = fmt.Sprintf("%s%s", warehouseHost, viper.GetString("client.warehouse.stockproduct"))
		userHost      = viper.GetString("client.user.host")
		urlUser       = fmt.Sprintf("%s%s", userHost, viper.GetString("client.user.stockproduct"))
	)
	tx := w.db.MustBegin()

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	}

	client := request.NewHTTPClient(w.httpClient)
	reqHttp := client.Request(header, urlUser+"?userId="+userAccountId)

	res, statusCode, err := reqHttp.Get(ctx)
	if err != nil {
		tx.Rollback()
		logger.Messagef("error request to warehouse service %v", err).To(ctx)
		err = fmt.Errorf("error request to warehouse service")
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Get Warehouse", err)
	}

	err = json.Unmarshal(res, &customer)
	if err != nil {
		tx.Rollback()
		err = fmt.Errorf("invalid umarshal json")
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Get Warehouse unmarshal", err)
	}

	for _, product := range *reqs {
		totalPrice += product.Price * float64(product.Quantity)
		TotalWeight += product.Weight * float64(product.Quantity)
		TotalQuantity += product.Quantity

	}
	order.CustomerId = int(customer.ID)
	order.IsPaymented = false
	order.IsDeleted = false
	order.TotalPrice = totalPrice
	order.TotalQuantity = TotalQuantity
	order.TotalWeight = TotalWeight
	idOrder, err := w.orderRepo.InsertOrUpdateOrder(tx, *order)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Create Order", err)
	}

	for _, product := range *reqs {
		orderProduct := models.OrderProduct{
			ProductId:   product.ProductID,
			OrderId:     idOrder,
			Price:       product.Price,
			WarehouseId: product.WarehouseId,
			Weight:      product.Weight,
			Quantity:    product.Quantity,
		}
		idProduct, err := w.orderRepo.InsertOrUpdateOrderProduct(tx, orderProduct)

		reqHttp = client.Request(header, urlWarehouse)
		stock := models.Stock{
			ProductId:   product.ProductID,
			WarehouseId: product.WarehouseId,
			Quantity:    product.Quantity,
			Type:        "decrease",
			IsDeleted:   false,
		}

		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()
			return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Create Order", err)
		}
		fmt.Println(idProduct)
		payload, err := json.Marshal(stock)
		if err != nil {
			err = fmt.Errorf("invalid marshal json")
			tx.Rollback()
			return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Create Order", err)
		}
		res, statusCode, err = reqHttp.Put(ctx, payload)
		if err != nil {
			tx.Rollback()
			logger.Messagef("error request to warehouse service %v", err).To(ctx)
			err = fmt.Errorf("error request to warehouse service")
			return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Get Warehouse", err)
		}

		err = json.Unmarshal(res, &customer)
		if err != nil {
			tx.Rollback()
			err = fmt.Errorf("invalid umarshal json")
		}

		if statusCode >= 400 {
			err = fmt.Errorf("response error")
			return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Create Warehouse unmarshal", err)
		}
	}

	err = tx.Commit()

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Order", "Failed Create Order", err)
	}

	return response.Success(ctx, http.StatusCreated, "000001", "Success Insert or Update Shop", nil)
}
