package postgresql

import (
	"context"
	"fmt"
	"main/internal/models"
)

func (r *Repository) DeleteStockByID(ctx context.Context, stock models.Stock) (err error) {
	var query = `
		DELETE FROM public.goods_stock
		WHERE stock_id = ?;

		DELETE FROM public.stock
		WHERE id = ?;
	`
	_, err = r.db.Exec(query, stock.ID, stock.ID)
	fmt.Printf("\n\nerr: %v\n\n", err)
	if err != nil {
		return err
	}
	return nil
}
