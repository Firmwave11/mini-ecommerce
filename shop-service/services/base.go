package services

import (
	"context"
	"crypto/tls"
	"net/http"
	models "shop-service/models/entity"
	shopRepo "shop-service/repositories"
	"shop-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type shopService struct {
	shopRepo   shopRepo.ShopRepo
	db         *sqlx.DB
	httpClient *http.Client
}

type ShopService interface {
	GetShop(ctx context.Context) response.Output
	InsertorUpdateShop(ctx context.Context, req models.Shop) response.Output
}

func NewShopService(e shopRepo.ShopRepo, db *sqlx.DB) ShopService {
	clientHttp := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &shopService{
		shopRepo:   e,
		db:         db,
		httpClient: clientHttp,
	}
}
