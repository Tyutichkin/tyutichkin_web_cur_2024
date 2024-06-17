package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) DeleteGoodByID(ctx context.Context, good models.Good) (err error) {
	var query = `
		DELETE FROM main.goods_stock
		WHERE goods_id = ?;

		DELETE FROM main.goods
		WHERE id = ?;
	`
	_, err = r.db.Exec(query, good.ID, good.ID)
	if err != nil {
		return err
	}
	return nil
}
