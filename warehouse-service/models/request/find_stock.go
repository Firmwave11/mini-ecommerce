package request

type GetStock struct {
	ProductId   int `json:"product_id"`
	WarehouseId int `json:"warhouse_id"`
}
