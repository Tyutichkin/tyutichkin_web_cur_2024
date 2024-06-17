package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) EditUser(ctx context.Context, user models.User) (err error) {
	var query = `
		UPDATE main.user SET
				fullname = ?,
				login = ?,
				password = ?,
				is_admin = ?
		WHERE id = ?;
	`
	_, err = r.db.Exec(query, user.FullName, user.Login, user.Password, user.IsAdmin, user.ID)
	if err != nil {
		return err
	}
	return nil
}
