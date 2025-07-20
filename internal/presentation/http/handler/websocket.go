package handler

import (
	"log/slog"
	"net/http"

	"github.com/K-Kizuku/ito-denwa/internal/application/service"
	"github.com/K-Kizuku/ito-denwa/internal/application/usecase"
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IWebSocketHandler interface {
	WebSocket(c *gin.Context)
	DebugWebSocket(c *gin.Context)
}

type WebSocketHandler struct {
	ItodenwaUsecase usecase.IItodenwaUsecase
}

func NewWebSocketHandler(ItodenwaUsecase usecase.IItodenwaUsecase) IWebSocketHandler {
	return &WebSocketHandler{
		ItodenwaUsecase: ItodenwaUsecase,
	}
}

func (h *WebSocketHandler) WebSocket(c *gin.Context) {
	roomID := c.Query("room_id")
	pc := c.Query("pc")
	if pc == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_id and pc are required"})
		return
	}
	isPC := pc == "true"
	slog.Info("WebSocket connection established", "room_id", roomID, "isPC", isPC)
	if roomID == "" {
		h.ItodenwaUsecase.CreateRoom(c.Request.Context())
	}
	conn, err := service.Upgrade(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	err = h.ItodenwaUsecase.AddPool(c.Request.Context(), entity.User{Tel: roomID}, *conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to pool"})
		return
	}
}

func (h *WebSocketHandler) DebugWebSocket(c *gin.Context) {
	conn, err := service.Upgrade(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	ctx := c.Request.Context()
	textCh := make(chan string)
	binaryCh := make(chan []byte)
	errCh := make(chan error)
	defer func() {
		close(textCh)
		close(binaryCh)
		close(errCh)
	}()

	go func() {
		for {
			conn.Read(ctx, textCh, binaryCh, errCh)
		}
	}()

	for {
		select {
		case text := <-textCh:
			if err := conn.Send(ctx, websocket.TextMessage, []byte(text)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send text message"})
				return
			}
		case binary := <-binaryCh:
			if err := conn.Send(ctx, websocket.BinaryMessage, binary); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send binary message"})
				return
			}
		case err := <-errCh:
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		case <-ctx.Done():
			c.JSON(http.StatusOK, gin.H{"status": "WebSocket connection closed"})
			return
		}
	}
}
