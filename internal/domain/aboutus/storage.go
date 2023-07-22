package aboutus

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wibedev-team/datachain-back/internal/models"
	"log"
)

type storage struct {
	repo *pgxpool.Pool
}

func NewStorage(r *pgxpool.Pool) *storage {
	return &storage{
		repo: r,
	}
}

func (s *storage) SaveSection(ctx context.Context, dto models.About) error {
	query := `
		INSERT INTO about (title, description, img)
		VALUES ($1, $2, $3)
	`

	exec, err := s.repo.Exec(ctx, query, dto.Title, dto.Description, dto.Img)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(exec.RowsAffected())

	return nil
}

func (s *storage) GetSection(ctx context.Context) (models.About, error) {
	query := `
		SELECT title, description, img 
		FROM about
		ORDER BY created_at DESC   
		LIMIT 1
	`

	var dto models.About
	err := s.repo.QueryRow(ctx, query).Scan(&dto.Title, &dto.Description, &dto.Img)
	if err != nil {
		log.Println(err)
		return models.About{}, err
	}

	return dto, nil
}
