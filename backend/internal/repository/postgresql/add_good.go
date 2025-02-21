package postgresql

import (
	"context"
	"main/internal/models"
	"time"
)

func (r *Repository) AddGood(ctx context.Context, good models.Good) (err error) {
	var query = `
		INSERT INTO public.goods(name, description, price, created_by_user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = r.db.Exec(query,
		good.Name,
		good.Description,
		good.Price,
		good.CreatedByUserID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
