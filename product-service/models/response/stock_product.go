package models

import models "product-service/models/entity"

type StatusData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type ResponseStock struct {
	Status *StatusData     `json:"status"`
	Stock  *[]models.Stock `json:"data"`
}
