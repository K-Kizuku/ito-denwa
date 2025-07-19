package handler

import (
	"net/http"

	"github.com/K-Kizuku/ito-denwa/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type IHealthzHandler interface {
	Healthz(c *gin.Context)
}

type HealthzHandler struct {
	HealthzUsecase usecase.IHealthzUsecase
}

func NewHealthzHandler(healthzUsecase usecase.IHealthzUsecase) IHealthzHandler {
	return &HealthzHandler{
		HealthzUsecase: healthzUsecase,
	}
}

func (h *HealthzHandler) Healthz(c *gin.Context) {
	healthz, err := h.HealthzUsecase.Healthz()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": healthz.Status})
}
