package postgresql

import (
	"context"
	"main/internal/models"
)

func (r *Repository) EditGood(ctx context.Context, good models.Good) (err error) {
	var query = `
		UPDATE public.goods SET
				name = ?,
				description = ?,
				price = ?
		WHERE id = ?;
	`
	_, err = r.db.Exec(query, good.Name, good.Description, good.Price, good.ID)
	if err != nil {
		return err
	}
	return nil
}
