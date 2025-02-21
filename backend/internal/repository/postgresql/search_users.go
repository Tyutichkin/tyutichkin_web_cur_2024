package postgresql

import (
	"context"
	"main/internal/models"
)

func (r *Repository) SearchUsers(ctx context.Context, searchRequest models.SearchUserRequest) (users []models.User, err error) {
	var query = `
		SELECT id,
		       fullname,
		       login,
		       password,
		       is_admin
		FROM public.user
		WHERE ($1 = '' OR (fullname like concat_ws('%', $1, '%')))
		AND ($2 = '' OR (login like concat_ws('%', $2, '%')))
		ORDER BY id;
	`
	rows, err := r.db.Query(query, searchRequest.FullName, searchRequest.Login)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Login,
			&user.Password,
			&user.IsAdmin,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
