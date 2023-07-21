package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wibedev-team/datachain-back/internal/models"
)

type storage struct {
	repo *pgxpool.Pool
}

func NewStorage(r *pgxpool.Pool) *storage {
	return &storage{
		repo: r,
	}
}

func (s *storage) FindUserByLogin(ctx context.Context, login string) (models.User, error) {
	query := `
		SELECT login, password, role 
		FROM users 
		WHERE login = $1
    `

	var user models.User
	err := s.repo.QueryRow(ctx, query, login).Scan(&user.Login, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
