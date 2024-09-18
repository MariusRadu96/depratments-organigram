package users

import (
	"context"
	"database/sql"
	"departments-organigram/internal/core/domain"
)

type store struct {
	db *sql.DB
}

func NewStore(database *sql.DB) *store {
	return &store{
		db: database,
	}
}

func (s *store) CreateUser(ctx context.Context, user domain.User) error {
	_, err := s.db.Query("CALL CreateUser(?, ?)", user.Username, user.Password)
	return err
}

func (s *store) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	rows, err := s.db.Query("CALL GetUserByUsername(?)", username)
	if err != nil {
		return domain.User{}, err
	}
	defer rows.Close()

	var user domain.User
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		); err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}
