package service

import (
	"context"
	"errors"
	"net/http"

	gorilla "github.com/gorilla/websocket"
)

type IWebsocket interface {
	Send(ctx context.Context, messageType int, data []byte) error
	SendBinary(ctx context.Context, ch <-chan []byte, errCh chan<- error)
	SendText(ctx context.Context, ch <-chan []byte, errCh chan<- error)
	Read(ctx context.Context, textCh chan<- string, binaryCh chan<- []byte, errCh chan<- error)
	Close() error
}

type Websocket struct {
	Conn *gorilla.Conn
}

func NewWebsocket(conn *gorilla.Conn) *Websocket {
	return &Websocket{
		Conn: conn,
	}
}

var _ IWebsocket = (*Websocket)(nil)

var upgrader = gorilla.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*Websocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	ws := NewWebsocket(conn)
	return ws, nil
}

func (ws *Websocket) Send(ctx context.Context, messageType int, data []byte) error {
	if err := ws.Conn.WriteMessage(messageType, data); err != nil {
		return err
	}
	return nil
}

func (ws *Websocket) SendBinary(ctx context.Context, ch <-chan []byte, errCh chan<- error) {
	for {
		select {
		case data := <-ch:
			if err := ws.Send(ctx, gorilla.BinaryMessage, data); err != nil {
				errCh <- err
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (ws *Websocket) SendText(ctx context.Context, ch <-chan []byte, errCh chan<- error) {
	for {
		select {
		case data := <-ch:
			if err := ws.Send(ctx, gorilla.TextMessage, data); err != nil {
				errCh <- err
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (ws *Websocket) Read(ctx context.Context, textCh chan<- string, bynaryCh chan<- []byte, errCh chan<- error) {
	i, data, err := ws.Conn.ReadMessage()
	if err != nil {
		errCh <- err
		return
	}
	switch i {
	case gorilla.TextMessage:
		select {
		case textCh <- string(data):
		case <-ctx.Done():
			return
		}
	case gorilla.BinaryMessage:
		select {
		case bynaryCh <- data:
		case <-ctx.Done():
			return
		}
	default:
		errCh <- errors.New("unsupported message type")
		return
	}
}

func (ws *Websocket) Close() error {
	if err := ws.Conn.Close(); err != nil {
		return err
	}
	return nil
}
