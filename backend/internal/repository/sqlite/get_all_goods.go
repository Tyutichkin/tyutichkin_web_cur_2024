package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) GetAllGoods(ctx context.Context) (goods []models.Good, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	defer tx.Rollback()
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
		FROM main.goods g
		LEFT JOIN main.user u ON u.id = g.created_by_user_id
		LEFT JOIN main.goods_stock gs on gs.goods_id = g.id
		ORDER BY price, gs.goods_count
	`
	rows, err := tx.Query(query)
	defer rows.Close()
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
