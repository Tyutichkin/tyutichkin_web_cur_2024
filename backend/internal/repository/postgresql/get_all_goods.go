package postgresql

import (
	"context"
	"database/sql"
	"log"
	"main/internal/models"
)

func (r *Repository) GetAllGoods(ctx context.Context) (goods []models.Good, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Print("failed to rollback transaction", err)
		}
	}(tx)
	if err != nil {
		return nil, err
	}
	var query = `
		SELECT g.id,
		       name,
		       description,
		       price,
		       coalesce(u.fullname, ''),
		       coalesce(gs.stock_id, 0),
		       coalesce(gs.goods_count, 0)
		FROM public.goods g
		LEFT JOIN public.user u ON u.id = g.created_by_user_id
		LEFT JOIN public.goods_stock gs on gs.goods_id = g.id
		ORDER BY price, gs.goods_count
	`
	rows, err := tx.Query(query)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Print("failed to close rows:", err)
		}
	}(rows)
	if err != nil {
		return []models.Good{}, err
	}

	for rows.Next() {
		var good models.Good
		err := rows.Scan(
			&good.ID,
			&good.Name,
			&good.Description,
			&good.Price,
			&good.CreatedByUserFullName,
			&good.StockID,
			&good.Count,
		)
		if err != nil {
			return nil, err
		}
		goods = append(goods, good)
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return goods, nil
}
