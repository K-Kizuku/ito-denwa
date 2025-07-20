//go:build wireinject
// +build wireinject

package di

import (
	"github.com/K-Kizuku/ito-denwa/internal/application/service"
	"github.com/K-Kizuku/ito-denwa/internal/application/usecase"
	"github.com/K-Kizuku/ito-denwa/internal/config"
	"github.com/K-Kizuku/ito-denwa/internal/infrastructure/repository"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/http"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/http/handler"
	"github.com/K-Kizuku/ito-denwa/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitRouter(e *gin.Engine, cfg *config.Config) *http.Router {
	wire.Build(
		repository.NewHealthzRepository,
		repository.NewRoomRepository,
		usecase.NewHealthzUsecase,
		usecase.NewItodenwaUsecase,
		handler.NewHealthzHandler,
		handler.NewWebSocketHandler,
		usecase.NewItodenwaUsecase,
		jwt.NewJWT,
		service.NewCardService,
		service.NewStringItemService,
		service.NewUserService,
		http.NewRouter,
	)
	return &http.Router{}
}
