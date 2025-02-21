package postgresql

import (
	"context"
	"main/internal/models"
)

func (r *Repository) DeleteUserByID(ctx context.Context, user models.User) (err error) {
	var query = `
		DELETE FROM public.user
		WHERE id = $1;
	`
	_, err = r.db.Exec(query, user.ID)
	if err != nil {
		return err
	}
	return nil
}
