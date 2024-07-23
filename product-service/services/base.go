package services

import (
	"context"
	"crypto/tls"
	"net/http"
	models "product-service/models/entity"
	productRepo "product-service/repositories"
	"product-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type productService struct {
	productRepo productRepo.ProductRepo
	db          *sqlx.DB
	httpClient  *http.Client
}

type ProductService interface {
	GetProduct(ctx context.Context, token string) response.Output
	InsertOrUpdateProduct(ctx context.Context, req models.Product) response.Output
}

func NewProductService(e productRepo.ProductRepo, db *sqlx.DB) ProductService {
	clientHttp := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return &productService{
		productRepo: e,
		db:          db,
		httpClient:  clientHttp,
	}
}
