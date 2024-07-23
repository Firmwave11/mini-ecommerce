package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"product-service/logger"
	models "product-service/models/entity"
	resModels "product-service/models/response"
	"product-service/utils/request"
	"product-service/utils/response"
	"strconv"

	"github.com/spf13/viper"
)

func (w *productService) GetProduct(ctx context.Context, token string) response.Output {

	var (
		warehouseHost = viper.GetString("client.warehouse.host")
		urlWarehouse  = fmt.Sprintf("%s%s", warehouseHost, viper.GetString("client.warehouse.stockproduct"))
		viewProducts  = []resModels.ProductResponse{}
		stocks        = resModels.ResponseStock{}
	)

	responseProduct, err := w.productRepo.FindAllProduct()
	if err != nil && err != sql.ErrNoRows {
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Get Product", "Failed Get Product", err)
	}

	queryParams := url.Values{}

	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": token,
	}

	for _, product := range responseProduct {
		queryParams.Add("productIds", strconv.Itoa(product.ID))
	}

	client := request.NewHTTPClient(w.httpClient)
	reqHttp := client.Request(header, urlWarehouse+"?"+queryParams.Encode())

	res, statusCode, err := reqHttp.Get(ctx)
	if err != nil {
		logger.Messagef("error request to warehouse service %v", err).To(ctx)
		err = fmt.Errorf("error request to warehouse service")
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Get Product", "Failed Get Warehouse", err)
	}

	err = json.Unmarshal(res, &stocks)
	if err != nil {
		err = fmt.Errorf("invalid umarshal json")
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Get Product", "Failed Get Warehouse unmarshal", err)
	}

	for _, product := range responseProduct {
		viewProduct := resModels.ProductResponse{}
		viewProduct.ID = product.ID
		viewProduct.Enabled = product.Enabled
		viewProduct.Name = product.Name
		viewProduct.Price = product.Price
		viewProduct.Weight = product.Weight
		viewProduct.CreatedAT = product.CreatedAT
		viewProduct.UpdatedAT = product.UpdatedAT
		for _, stock := range *stocks.Stock {
			if stock.ProductId == product.ID {
				viewProduct.Stock = append(viewProduct.Stock, &stock)
			}
		}
		viewProducts = append(viewProducts, viewProduct)
	}

	if statusCode >= 400 {
		err = fmt.Errorf("response error")
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Get Product", "Failed Get Warehouse unmarshal", err)
	}

	return response.Success(ctx, http.StatusOK, "000001", "Success Get Product", viewProducts)
}

func (w *productService) InsertOrUpdateProduct(ctx context.Context, req models.Product) response.Output {
	tx := w.db.MustBegin()

	err := w.productRepo.InsertOrUpdateProduct(tx, req)
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Insert or Update Product", "Failed Insert or Update Product", err)
	}

	err = tx.Commit()

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Insert or Update Product", "Failed Insert or Update Product", err)
	}

	return response.Success(ctx, http.StatusCreated, "000002", "Success Insert or Update Product", nil)
}
