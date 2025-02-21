package postgresql

import (
	"context"
	"main/internal/models"
)

func (r *Repository) AddStock(ctx context.Context, stock models.Stock) (err error) {
	var query = `
		INSERT INTO public.stock(adress)
		VALUES ($1)
	`
	_, err = r.db.Exec(query, stock.Address)
	if err != nil {
		return err
	}
	return nil
}
