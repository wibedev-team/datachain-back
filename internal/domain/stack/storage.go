package stack

import (
	"context"
	"log"

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

func (s *storage) CreateStackImage(ctx context.Context, img string) error {
	query := `
		INSERT INTO stacks (img)
		VALUES ($1)
    `

	exec, err := s.repo.Exec(ctx, query, img)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(exec.RowsAffected())
	return nil
}

func (s *storage) GetAllStackImages(ctx context.Context) ([]models.Stack, error) {
	query := `
		SELECT img 
		FROM stacks 
    `

	rows, err := s.repo.Query(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var stacks []models.Stack
	for rows.Next() {
		var img models.Stack
		err := rows.Scan(&img.Img)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		stacks = append(stacks, img)
	}

	return stacks, nil
}

func (s *storage) RemoveStack(ctx context.Context, id string) error {
	query := `
		DELETE FROM stacks
		WHERE img = $1
	`

	exec, err := s.repo.Exec(ctx, query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(exec.RowsAffected())

	return nil
}
