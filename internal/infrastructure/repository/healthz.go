package repository

import (
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/internal/domain/repository"
)

type HealthzRepository struct{}

func NewHealthzRepository() repository.IHealthzRepository {
	return &HealthzRepository{}
}

func (r *HealthzRepository) Healthz() (*entity.Healthz, error) {
	return &entity.Healthz{
		Status: "OK",
	}, nil
}
