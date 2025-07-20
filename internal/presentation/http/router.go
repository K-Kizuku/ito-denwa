package http

import (
	"log/slog"
	"os"

	"github.com/K-Kizuku/ito-denwa/internal/config"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/http/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	hh  handler.IHealthzHandler
	wsh handler.IWebSocketHandler
}

type IRouter interface {
	Setup(e *gin.Engine, cfg *config.Config) *gin.Engine
}

var _ IRouter = (*Router)(nil)

func NewRouter(hh handler.IHealthzHandler, wsh handler.IWebSocketHandler) *Router {
	return &Router{
		hh:  hh,
		wsh: wsh,
	}
}

func (r *Router) Setup(e *gin.Engine, cfg *config.Config) *gin.Engine {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	api := e.Group("/api")
	{
		api.GET("/healthz", r.hh.Healthz)

		ws := api.Group("/ws")
		{
			ws.GET("/pc", r.wsh.WebSocketPC)
			ws.GET("/mobile", r.wsh.WebSocketMobile)
			ws.GET("/debug", r.wsh.DebugWebSocket)
		}
	}
	return e
}
