package repository

import (
	"replace_go/model/auth"
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) GetUserID(input auth.SignIn) (*auth.AuthResponse, error) {
	query := `
		SELECT id
		FROM users
		WHERE auth_id = $1
	`
	var userID int
	err := r.DB.QueryRow(query, input.AuthId).Scan(&input.AuthId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return &auth.AuthResponse{Id: userID}, nil
}