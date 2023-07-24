package solutions

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

func (s *storage) SaveSolution(ctx context.Context, dto models.Solution) error {
	query := `
		INSERT INTO solutions (title, link, file)
		VALUES ($1, $2, $3)
	`

	exec, err := s.repo.Exec(ctx, query, dto.Title, dto.Link, dto.File)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(exec.RowsAffected())

	query = `
		INSERT INTO features (title, descr)
		VALUES ($1, $2)
	`
	for _, feature := range dto.Features {
		exec, err = s.repo.Exec(ctx, query, dto.Title, feature.Text)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(exec.RowsAffected())
	}

	return nil
}

func (s *storage) GetAllSolutions(ctx context.Context) ([]models.Solution, error) {
	query := `
		SELECT title, link, file  
		FROM solutions
	`

	rows, err := s.repo.Query(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var solutions []models.Solution
	for rows.Next() {
		var sol models.Solution
		err := rows.Scan(&sol.Title, &sol.Link, &sol.File)
		if err != nil {
			log.Println(err)
		}

		query := `
			SELECT descr  
			FROM features
			WHERE title = $1
		`
		r, err := s.repo.Query(ctx, query, sol.Title)
		if err != nil {
			log.Println(err)
			continue
		}
		var features []models.Feature
		for r.Next() {
			var f models.Feature
			err := r.Scan(&f.Text)
			if err != nil {
				log.Println(err)
				continue
			}

			features = append(features, f)
		}

		sol.Features = features
		solutions = append(solutions, sol)
	}

	return solutions, nil
}

func (s *storage) RemoveSolution(ctx context.Context, title string) error {
	query := `
		DELETE FROM features
		WHERE title = $1
	`
	exec, err := s.repo.Exec(ctx, query, title)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(exec.RowsAffected())

	query = `
		DELETE FROM solutions
		WHERE title = $1
	`
	exec, err = s.repo.Exec(ctx, query, title)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(exec.RowsAffected())

	return nil
}
