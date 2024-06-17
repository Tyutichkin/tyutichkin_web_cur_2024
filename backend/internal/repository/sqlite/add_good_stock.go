package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) AddGoodStock(ctx context.Context, goodStock models.GoodStock) (err error) {
	var query = `
		INSERT INTO main.goods_stock(goods_id, stock_id, goods_count)
		VALUES ($1, $2, $3)
	`
	_, err = r.db.Exec(query, goodStock.GoodID, goodStock.StockID, goodStock.GoodCount)
	if err != nil {
		return err
	}
	return nil
}
