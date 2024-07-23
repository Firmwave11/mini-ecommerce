package repositories

import (
	"order-service/middleware"
	models "order-service/models/entity"

	"github.com/jmoiron/sqlx"
)

var timezone = middleware.Timezone()

type orderRepo struct {
	db *sqlx.DB
}

type OrderRepo interface {
	InsertOrUpdateOrderProduct(tx *sqlx.Tx, orderProductReq models.OrderProduct) (int, error)
	InsertOrUpdateOrder(tx *sqlx.Tx, orderReq models.Order) (int, error)
	FindAllOrder() ([]*models.Order, error)
}

func NewOrderRepository(db *sqlx.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}
