package services

import (
	"context"
	"crypto/tls"
	"net/http"
	req "order-service/models/request"
	orderRepo "order-service/repositories"
	"order-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type orderService struct {
	orderRepo  orderRepo.OrderRepo
	db         *sqlx.DB
	httpClient *http.Client
}

type OrderService interface {
	CreateOrder(ctx context.Context, reqs *[]req.ProductRequest, userAccountId string, token string) response.Output
}

func NewOrderService(e orderRepo.OrderRepo, db *sqlx.DB) OrderService {
	clientHttp := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &orderService{
		orderRepo:  e,
		db:         db,
		httpClient: clientHttp,
	}
}
