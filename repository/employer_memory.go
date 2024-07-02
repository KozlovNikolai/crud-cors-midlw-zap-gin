package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/models"
)

type InMemoryEmployerRepository struct {
	employers map[int]models.Employer
	nextID    int
	mutex     sync.Mutex
}

func NewInMemoryEmployerRepository() *InMemoryEmployerRepository {
	return &InMemoryEmployerRepository{
		employers: make(map[int]models.Employer),
		nextID:    1,
	}
}

func (repo *InMemoryEmployerRepository) CreateEmployer(ctx context.Context, employer models.Employer) (int, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	employer.ID = repo.nextID
	repo.nextID++
	repo.employers[employer.ID] = employer
	return employer.ID, nil
}

func (repo *InMemoryEmployerRepository) GetEmployerByID(ctx context.Context, id int) (models.Employer, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	employer, exists := repo.employers[id]
	if !exists {
		return employer, errors.New("employer not found")
	}
	return employer, nil
}

func (repo *InMemoryEmployerRepository) GetAllEmployers(ctx context.Context) ([]models.Employer, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	var employers []models.Employer
	for _, employer := range repo.employers {
		employers = append(employers, employer)
	}
	return employers, nil
}

func (repo *InMemoryEmployerRepository) UpdateEmployer(ctx context.Context, id int, employer models.Employer) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if _, exists := repo.employers[id]; !exists {
		return errors.New("employer not found")
	}
	employer.ID = id
	repo.employers[id] = employer
	return nil
}

func (repo *InMemoryEmployerRepository) DeleteEmployer(ctx context.Context, id int) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	if _, exists := repo.employers[id]; !exists {
		return errors.New("employer not found")
	}
	delete(repo.employers, id)
	return nil
}
