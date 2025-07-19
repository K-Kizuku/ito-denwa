package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/K-Kizuku/ito-denwa/internal/config"
	"github.com/K-Kizuku/ito-denwa/internal/di"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	cfg := config.GetHttpConfig()
	router := di.InitRouter(engine, cfg)
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router.Setup(engine, cfg),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	<-ctx.Done()
	log.Println("Server exiting")
}
