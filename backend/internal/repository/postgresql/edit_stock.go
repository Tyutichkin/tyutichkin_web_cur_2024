package postgresql

import (
	"context"
	"main/internal/models"
)

func (r *Repository) EditStock(ctx context.Context, stock models.Stock) (err error) {
	var query = `
		UPDATE public.stock SET
				adress = ?
		WHERE id = ?;
	`
	_, err = r.db.Exec(query, stock.Address, stock.ID)
	if err != nil {
		return err
	}
	return nil
}
