package team

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

func (s *storage) SaveTeammate(ctx context.Context, dto models.Team) error {
	query := `
		INSERT INTO team (name, position, link, img)
		VALUES ($1, $2, $3, $4)
	`

	exec, err := s.repo.Exec(ctx, query, dto.Name, dto.Position, dto.Link, dto.Img)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(exec.RowsAffected())

	return nil
}

func (s *storage) GetAllTeammates(ctx context.Context) ([]models.Team, error) {
	query := `
		SELECT name, position, link, img 
		FROM team
	`

	rows, err := s.repo.Query(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var teammates []models.Team
	for rows.Next() {
		var t models.Team
		err := rows.Scan(&t.Name, &t.Position, &t.Link, &t.Img)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		teammates = append(teammates, t)
	}

	return teammates, nil
}

func (s *storage) RemoveTeammate(ctx context.Context, id string) error {
	query := `
		DELETE FROM team
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
