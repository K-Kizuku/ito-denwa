package repository

import (
	"sync"

	"github.com/cockroachdb/errors"

	"github.com/K-Kizuku/ito-denwa/internal/application/service"
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/pkg/null"
)

type Room struct {
	Room        entity.Room
	Connections Connection
	BinaryCh    chan []byte
	JSONCh      chan string
	ErrCh       chan error
}

type Connection struct {
	CallerPC     null.Null[service.Websocket]
	CallerMobile null.Null[service.Websocket]
	CalleePC     null.Null[service.Websocket]
	CalleeMobile null.Null[service.Websocket]
}

type RoomRepository struct {
	mu    sync.RWMutex
	rooms map[string]Room
}

type IRoomRepository interface {
	Create(room entity.Room) error
	AddConnection(roomID string, user *entity.User, conn *service.Websocket, isCaller, isPC bool) error
	AddUser(roomID string, user entity.User, isCaller bool) error
	RemoveUser(roomID string, userID string) error
	GetRoom(roomID string) (*Room, error)
	GetChannels(roomID string) (chan []byte, chan string, chan error, error)
}

func NewRoomRepository() IRoomRepository {
	return &RoomRepository{
		rooms: make(map[string]Room),
	}
}

func (r *RoomRepository) Create(room entity.Room) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.rooms[room.ID]; exists {
		return errors.New("room already exists")
	}

	newRoom := &Room{
		Room: room,
		Connections: Connection{
			CallerPC:     null.New[service.Websocket](nil),
			CallerMobile: null.New[service.Websocket](nil),
			CalleePC:     null.New[service.Websocket](nil),
			CalleeMobile: null.New[service.Websocket](nil),
		},
		BinaryCh: make(chan []byte, 100),
		JSONCh:   make(chan string, 100),
		ErrCh:    make(chan error, 100),
	}

	r.rooms[room.ID] = *newRoom
	return nil
}

func (r *RoomRepository) AddConnection(roomID string, user *entity.User, conn *service.Websocket, isCaller, isPC bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	room, exists := r.rooms[roomID]
	if !exists {
		return errors.New("room not found")
	}

	if isCaller {
		if isPC {
			room.Connections.CallerPC = null.New(conn)
		} else {
			room.Connections.CallerMobile = null.New(conn)
		}
	} else {
		if isPC {
			room.Connections.CalleePC = null.New(conn)
		} else {
			room.Connections.CalleeMobile = null.New(conn)
		}
	}

	r.rooms[roomID] = room
	return nil
}

func (r *RoomRepository) AddUser(roomID string, user entity.User, isCaller bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	room, exists := r.rooms[roomID]
	if !exists {
		return errors.New("room not found")
	}
	if isCaller {
		room.Room.Caller = user
	} else {
		room.Room.Callee = user
	}
	r.rooms[roomID] = room
	return nil
}

func (r *RoomRepository) RemoveUser(roomID string, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	room, exists := r.rooms[roomID]
	if !exists {
		return errors.New("room not found")
	}

	if room.Room.Caller.ID == userID {
		room.Room.Caller = entity.User{}
		room.Connections.CallerPC = null.New[service.Websocket](nil)
		room.Connections.CallerMobile = null.New[service.Websocket](nil)
	}
	if room.Room.Callee.ID == userID {
		room.Room.Callee = entity.User{}
		room.Connections.CalleePC = null.New[service.Websocket](nil)
		room.Connections.CalleeMobile = null.New[service.Websocket](nil)
	}
	r.rooms[roomID] = room
	return nil
}

func (r *RoomRepository) GetRoom(roomID string) (*Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, exists := r.rooms[roomID]
	if !exists {
		return nil, errors.New("room not found")
	}

	return &room, nil
}

func (r *RoomRepository) GetChannels(roomID string) (chan []byte, chan string, chan error, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, exists := r.rooms[roomID]
	if !exists {
		return nil, nil, nil, errors.New("room not found")
	}

	return room.BinaryCh, room.JSONCh, room.ErrCh, nil
}
