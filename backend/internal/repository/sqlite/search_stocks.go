package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) SearchStocks(ctx context.Context, searchRequest models.SearchStockRequest) (stocks []models.Stock, err error) {
	var query = `
		SELECT id,
		       adress
		FROM main.stock s
		WHERE ($1 = '' OR (adress like concat_ws('%', $1, '%')))
		ORDER BY id;
	`
	rows, err := r.db.Query(query, searchRequest.Address)
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
	return stocks, nil
}
