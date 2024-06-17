package models

import (
	"time"
)

type Good struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	Price                 int       `json:"price"`
	Count                 int       `json:"count"`
	StockID               int       `json:"stock_id"`
	CreatedByUserID       int       `json:"created_by_user_id"`
	CreatedByUserFullName string    `json:"created_by_user_full_name"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type SearchGoodRequest struct {
	ID          int
	MinPrice    int    `json:"min_price"`
	MaxPrice    int    `json:"max_price"`
	MinCount    int    `json:"min_count"`
	MaxCount    int    `json:"max_count"`
	Name        string `json:"name"`
	IsPriceDesc bool   `json:"is_price_desc"`
	IsCountDesc bool   `json:"is_count_desc"`
}
