package footer

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

func (s *storage) CreateFooter(ctx context.Context, footer models.Footer) error {
	query := `
		INSERT INTO footer (email, telephone, address)
		VALUES ($1, $2, $3)
    `

	exec, err := s.repo.Exec(ctx, query, footer.Email, footer.Telephone, footer.Address)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(exec.RowsAffected())
	return nil
}

func (s *storage) GetFooter(ctx context.Context) (models.Footer, error) {
	query := `
		SELECT email, telephone, address 
		FROM footer
		ORDER BY created_at DESC   
		LIMIT 1
    `

	var dto models.Footer
	err := s.repo.QueryRow(ctx, query).Scan(&dto.Email, &dto.Telephone, &dto.Address)
	if err != nil {
		log.Println(err)
		return models.Footer{}, err
	}

	return dto, nil
}
