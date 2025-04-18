package postgresql

import (
	"context"
	"database/sql"
	"log/slog"
	"main/internal/models"
)

func (r *Repository) GetAllStocks(ctx context.Context) (stocks []models.Stock, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			slog.Error("failed to rollback transaction", "err", err)
		}
	}(tx)
	if err != nil {
		return nil, err
	}
	var query = `
		SELECT id,
		       adress
		FROM public.stock s
		ORDER BY id
	`
	rows, err := tx.Query(query)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			slog.Error("failed to close rows", "err", err)
		}
	}(rows)
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
