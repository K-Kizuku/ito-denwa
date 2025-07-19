//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/ito-denwa/internal/application/usecase"
	"github.com/K-Kizuku/ito-denwa/internal/config"
	"github.com/K-Kizuku/ito-denwa/internal/infrastructure/repository"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/http"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/http/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitRouter(e *gin.Engine, cfg *config.Config) *http.Router {
	wire.Build(
		repository.NewHealthzRepository,
		usecase.NewHealthzUsecase,
		http.NewRouter,
		handler.NewHealthzHandler,
		handler.NewWebSocketHandler,
	)
	return &http.Router{}
}
