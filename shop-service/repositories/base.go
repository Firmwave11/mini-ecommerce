package repositories

import (
	"shop-service/middleware"
	models "shop-service/models/entity"

	"github.com/jmoiron/sqlx"
)

var timezone = middleware.Timezone()

type shopRepo struct {
	db *sqlx.DB
}

type ShopRepo interface {
	InsertOrUpdateShop(tx *sqlx.Tx, shopReq models.Shop) error
	FindAllShop() ([]*models.Shop, error)
}

func NewShopRepository(db *sqlx.DB) ShopRepo {
	return &shopRepo{
		db: db,
	}
}
