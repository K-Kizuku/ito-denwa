package repository

import (
	"sync"

	"github.com/K-Kizuku/ito-denwa/internal/application/service"
	"github.com/K-Kizuku/ito-denwa/pkg/null"
)

type Connection struct {
	CallerPC null.Null[service.Websocket]
	CallerSM null.Null[service.Websocket]
	CalleePC null.Null[service.Websocket]
	CalleeSM null.Null[service.Websocket]
}

type RoomRepository struct {
	mu    sync.RWMutex
	rooms map[string]Room
}
