package request

type UpdateSyncStockWarehouseRequest struct {
	Quantity        int `json:"quantity"`
	ProductId       int `json:"product_id"`
	FromWarehouseId int `json:"from_warhouse_id"`
	ToWarehouseId   int `json:"to_warhouse_id"`
}
