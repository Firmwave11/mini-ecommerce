package request

type ProductRequest struct {
	ProductID   int     `json:"product_id"`
	WarehouseId int     `json:"warehouse_id"`
	Quantity    int     `json:"quantity"`
	Weight      float64 `json:"weight"`
	Price       float64 `json:"price"`
}
