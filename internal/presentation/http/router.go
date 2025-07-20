package http

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/K-Kizuku/ito-denwa/internal/config"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/http/handler"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/cards/cardsconnect"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/strings/stringsconnect"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/user/userconnect"
	"github.com/K-Kizuku/ito-denwa/internal/application/service"

	"github.com/gin-gonic/gin"
)

type Router struct {
	hh          handler.IHealthzHandler
	wsh         handler.IWebSocketHandler
	cardService *service.CardService
	stringService *service.StringItemService
	userService *service.UserService
}

type IRouter interface {
	Setup(e *gin.Engine, cfg *config.Config) *gin.Engine
}

var _ IRouter = (*Router)(nil)

func NewRouter(hh handler.IHealthzHandler, wsh handler.IWebSocketHandler, cardService *service.CardService, stringService *service.StringItemService, userService *service.UserService) *Router {
	return &Router{
		hh:          hh,
		wsh:         wsh,
		cardService: cardService,
		stringService: stringService,
		userService: userService,
	}
}

func (r *Router) Setup(e *gin.Engine, cfg *config.Config) *gin.Engine {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(r.corsMiddleware(cfg))
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

	// Connect handlers
	cardPath, cardHandler := cardsconnect.NewCardServiceHandler(r.cardService)
	stringPath, stringHandler := stringsconnect.NewStringItemServiceHandler(r.stringService)
	userPath, userHandler := userconnect.NewUserServiceHandler(r.userService)

	e.Any(cardPath+"*any", gin.WrapH(cardHandler))
	e.Any(stringPath+"*any", gin.WrapH(stringHandler))
	e.Any(userPath+"*any", gin.WrapH(userHandler))
	return e
}

func (r *Router) corsMiddleware(cfg *config.Config) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.Server.AllowOrigin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Connect-Protocol-Version, Connect-Timeout-Ms")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})
}
