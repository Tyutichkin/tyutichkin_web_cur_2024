package models

type Stock struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
}

type SearchStockRequest struct {
	Address string `json:"address"`
}
