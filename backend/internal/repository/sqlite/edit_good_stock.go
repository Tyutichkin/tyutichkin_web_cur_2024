package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) EditGoodStock(ctx context.Context, goodStock models.GoodStock) (err error) {
	// своеобразный вариант upsert из-за ограничений sqlite
	var query = `
		INSERT INTO main.goods_stock(goods_id, stock_id, goods_count)
		SELECT ?, ?, ?
		WHERE NOT EXISTS (
		    SELECT 1
		    FROM main.goods_stock gs
		    WHERE gs.stock_id = ? AND gs.goods_id = ?
		);

		UPDATE main.goods_stock SET
				goods_count = ?
		WHERE stock_id = ? AND goods_id = ?;
	`
	_, err = r.db.Exec(query, goodStock.GoodID, goodStock.StockID, goodStock.GoodCount, goodStock.StockID, goodStock.GoodID, goodStock.GoodCount, goodStock.StockID, goodStock.GoodID)
	if err != nil {
		return err
	}
	return nil
}
