package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) GetAllStocks(ctx context.Context) (stocks []models.Stock, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	var query = `
		SELECT id,
		       adress
		FROM main.stock s
		ORDER BY id
	`
	rows, err := tx.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(
			&stock.ID,
			&stock.Address,
		)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return stocks, nil
}
