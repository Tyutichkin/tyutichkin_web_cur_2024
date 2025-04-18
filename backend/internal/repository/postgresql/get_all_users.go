package postgresql

import (
	"context"
	"database/sql"
	"log"
	"main/internal/models"
)

func (r *Repository) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Print("failed to rollback transaction", "err", err)
		}
	}(tx)
	if err != nil {
		return nil, err
	}
	var query = `
		SELECT id,
		       fullname,
		       login,
		       password,
		       is_admin
		FROM public.user
		ORDER BY id;
	`
	rows, err := tx.Query(query)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Print("failed to close rows", "err", err)
		}
	}(rows)
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
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return users, nil
}
