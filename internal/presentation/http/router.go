package http

import (
	"log/slog"
	"os"

	"github.com/K-Kizuku/ito-denwa/internal/config"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/http/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	hh handler.IHealthzHandler
}

type IRouter interface {
	Setup(e *gin.Engine, cfg *config.Config) *gin.Engine
}

var _ IRouter = (*Router)(nil)

func NewRouter(hh handler.IHealthzHandler) *Router {
	return &Router{
		hh: hh,
	}
}

func (r *Router) Setup(e *gin.Engine, cfg *config.Config) *gin.Engine {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	api := e.Group("/api")
	api.GET("/healthz", r.hh.Healthz)
	return e
}
