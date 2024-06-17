package sqlite

import (
	"context"
	"main/internal/models"
)

func (r *Repository) GetUserByLogin(ctx context.Context, login string) (user models.User, err error) {
	var query = `
		SELECT id,
		       fullname,
		       login
		FROM main.user
		WHERE login = $1;
	`
	err = r.db.QueryRow(query, login).Scan(
		&user.ID,
		&user.FullName,
		&user.Login)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
