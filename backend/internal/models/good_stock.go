package models

type GoodStock struct {
	GoodID    int `json:"good_id"`
	StockID   int `json:"stock_id"`
	GoodCount int `json:"good_count"`
}
