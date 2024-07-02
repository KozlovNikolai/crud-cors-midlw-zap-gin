package repository

import (
	"context"

	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/models"
)

type EmployerRepository interface {
	CreateEmployer(ctx context.Context, employer models.Employer) (int, error)
	GetEmployerByID(ctx context.Context, id int) (models.Employer, error)
	GetAllEmployers(ctx context.Context) ([]models.Employer, error)
	UpdateEmployer(ctx context.Context, id int, employer models.Employer) error
	DeleteEmployer(ctx context.Context, id int) error
}
