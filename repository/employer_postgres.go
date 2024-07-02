package repository

import (
	"context"

	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresEmployerRepository struct {
	db *pgxpool.Pool
}

func NewPostgresEmployerRepository(db *pgxpool.Pool) *PostgresEmployerRepository {
	return &PostgresEmployerRepository{db: db}
}

func (repo *PostgresEmployerRepository) CreateEmployer(ctx context.Context, employer models.Employer) (int, error) {
	var id int
	query := `
		INSERT INTO employers (company, name, email) 
		VALUES ($1, $2, $3) 
		RETURNING id`
	err := repo.db.QueryRow(ctx, query, employer.Company, employer.Person.Name, employer.Person.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *PostgresEmployerRepository) GetEmployerByID(ctx context.Context, id int) (models.Employer, error) {
	var employer models.Employer
	query := `
		SELECT e.id, e.company, p.id, p.name, p.email 
		FROM employers e 
		JOIN persons p ON e.person_id = p.id 
		WHERE e.id=$1`
	row := repo.db.QueryRow(ctx, query, id)
	err := row.Scan(&employer.ID, &employer.Company, &employer.Person.ID, &employer.Person.Name, &employer.Person.Email)
	if err != nil {
		return employer, err
	}
	return employer, nil
}

func (repo *PostgresEmployerRepository) GetAllEmployers(ctx context.Context) ([]models.Employer, error) {
	var employers []models.Employer
	query := `
		SELECT e.id, e.company, p.id, p.name, p.email 
		FROM employers e 
		JOIN persons p ON e.person_id = p.id`
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employer models.Employer
		err := rows.Scan(&employer.ID, &employer.Company, &employer.Person.ID, &employer.Person.Name, &employer.Person.Email)
		if err != nil {
			return nil, err
		}
		employers = append(employers, employer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return employers, nil
}

func (repo *PostgresEmployerRepository) UpdateEmployer(ctx context.Context, id int, employer models.Employer) error {
	query := `
		UPDATE employers 
		SET company=$1, name=$2, email=$3 
		WHERE id=$4`
	_, err := repo.db.Exec(ctx, query, employer.Company, employer.Person.Name, employer.Person.Email, id)
	return err
}

func (repo *PostgresEmployerRepository) DeleteEmployer(ctx context.Context, id int) error {
	query := "DELETE FROM employers WHERE id=$1"
	_, err := repo.db.Exec(ctx, query, id)
	return err
}
