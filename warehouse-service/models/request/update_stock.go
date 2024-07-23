package request

type UpdateStockRequest struct {
	Quantity    int    `json:"quantity"`
	ProductId   int    `json:"product_id"`
	WarehouseId int    `json:"warhouse_id"`
	Type        string `json:"type"`
}
