package postgresql

import (
	"context"
	"fmt"
	"main/internal/models"
)

func (r *Repository) SearchGoods(ctx context.Context, searchRequest models.SearchGoodRequest) (goods []models.Good, err error) {
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
		WHERE ($1 = '' OR (name like concat_ws('%', $1, '%')))
		AND ($2 = 0 OR (price >= $2))
		AND ($3 = 0 OR (price <= $3))
		AND ($4 = 0 OR (gs.goods_count >= $4))
		AND ($5 = 0 OR (gs.goods_count <= $5))
		AND ($6 = 0 OR (g.id = $6))
	`
	orderPriceStr := ""
	orderCountStr := ""
	if searchRequest.IsCountDesc == true {
		orderCountStr = "DESC"
	}
	if searchRequest.IsPriceDesc == true {
		orderPriceStr = "DESC"
	}
	query = fmt.Sprintf("%v\nORDER BY price %v, gs.goods_count %v", query, orderPriceStr, orderCountStr)
	rows, err := r.db.Query(query,
		searchRequest.Name,
		searchRequest.MinPrice,
		searchRequest.MaxPrice,
		searchRequest.MinCount,
		searchRequest.MaxCount,
		searchRequest.ID,
	)
	if err != nil {
		return nil, err
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
	return goods, nil
}
