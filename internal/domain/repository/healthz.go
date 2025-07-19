package repository

import "github.com/K-Kizuku/ito-denwa/internal/domain/entity"

type IHealthzRepository interface {
	Healthz() (*entity.Healthz, error)
}
