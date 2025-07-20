package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/K-Kizuku/ito-denwa/internal/application/service"
	"github.com/K-Kizuku/ito-denwa/internal/application/usecase"
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/pkg/uuid"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IWebSocketHandler interface {
	WebSocketMobile(c *gin.Context)
	WebSocketPC(c *gin.Context)
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

func (h *WebSocketHandler) WebSocketMobile(c *gin.Context) {
	tel := c.Query("tel")
	if tel == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tel is required"})
		return
	}
	slog.Info("WebSocket connection established for mobile", "tel", tel)
	conn, err := service.Upgrade(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	uid := uuid.New(c.Request.Context())
	err = h.ItodenwaUsecase.AddPool(c.Request.Context(), entity.User{ID: uid, Tel: tel}, *conn, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to pool"})
		return
	}
	time.Sleep(15 * time.Second) // Wait for the pool to be updated
	// 共通処理へ
	err = h.ItodenwaUsecase.Calling(c.Request.Context(), entity.User{ID: uid, Tel: tel}, *conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start calling"})
		return
	}
	slog.Info("WebSocket connection established for mobile", "tel", tel, "uid", uid)
}

func (h *WebSocketHandler) WebSocketPC(c *gin.Context) {
	myTel := c.Query("my_tel")
	targetTel := c.Query("target_tel")
	if targetTel == "" || myTel == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "my_tel and target_tel are required"})
		return
	}
	slog.Info("WebSocket connection established for PC", "my_tel", myTel, "target_tel", targetTel)
	conn, err := service.Upgrade(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	defer conn.Close()

	uid := uuid.New(c.Request.Context())
	err = h.ItodenwaUsecase.AddPool(c.Request.Context(), entity.User{ID: uid, Tel: myTel}, *conn, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to pool"})
		return
	}
	err = h.ItodenwaUsecase.CreateRoom(c.Request.Context(), myTel, targetTel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}
	err = h.ItodenwaUsecase.Calling(c.Request.Context(), entity.User{ID: uid, Tel: myTel}, *conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start calling"})
		return
	}
	slog.Info("WebSocket connection established for PC", "my_tel", myTel, "target_tel", targetTel, "uid", uid)
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
