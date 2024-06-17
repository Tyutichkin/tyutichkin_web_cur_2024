package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) AddUser(ctx context.Context, user models.User) (err error) {
	var query = `
		INSERT INTO main.user(fullname, login, password, is_admin)
		VALUES ($1, $2, $3, $4)
	`
	_, err = r.db.Exec(query, user.FullName, user.Login, user.Password, user.IsAdmin)
	if err != nil {
		return err
	}
	return nil
}
