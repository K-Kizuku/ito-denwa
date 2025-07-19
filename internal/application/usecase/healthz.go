package usecase

import (
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/internal/domain/repository"
)

type IHealthzUsecase interface {
	Healthz() (*entity.Healthz, error)
}

type HealthzUsecase struct {
	healthzRepository repository.IHealthzRepository
}

func NewHealthzUsecase(healthzRepository repository.IHealthzRepository) IHealthzUsecase {
	return &HealthzUsecase{
		healthzRepository: healthzRepository,
	}
}

func (u *HealthzUsecase) Healthz() (*entity.Healthz, error) {
	healthz, err := u.healthzRepository.Healthz()
	if err != nil {
		return nil, err
	}
	return healthz, nil
}
