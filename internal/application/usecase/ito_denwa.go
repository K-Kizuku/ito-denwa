package usecase

import (
	"context"
	"errors"
	"log/slog"
	"sync"

	"github.com/K-Kizuku/ito-denwa/internal/application/service"
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/internal/infrastructure/repository"
	"github.com/K-Kizuku/ito-denwa/pkg/uuid"
	"github.com/gorilla/websocket"
)

type IItodenwaUsecase interface {
	Receive(ctx context.Context, conn service.Websocket) error
	Send(ctx context.Context, conn service.Websocket) error
	Processing() error
	Calling(context.Context, entity.User, service.Websocket) error
	AddPool(ctx context.Context, user entity.User, conn service.Websocket, isPC bool) error
	RemovePool(ctx context.Context, user entity.User) error
	GetPool(ctx context.Context, tel string, isPC bool) *service.Websocket
	CreateRoom(ctx context.Context, callerTel, calleeTell string) error
	GetUser(context.Context, string, bool) *entity.User

	UseTelephoneCard() error
}

type ItodenwaUsecase struct {
	roomRepository repository.IRoomRepository
	Pool           map[string]PoolUserConn
	UserMap        map[string]PoolUser
	RoomMap        map[string]string //tel→roomID
	mu             sync.RWMutex
}

// Processing implements IItodenwaUsecase.
func (u *ItodenwaUsecase) Processing() error {
	panic("unimplemented")
}

// Send implements IItodenwaUsecase.
func (u *ItodenwaUsecase) Send(ctx context.Context, conn service.Websocket) error {
	panic("unimplemented")
}

// UseTelephoneCard implements IItodenwaUsecase.
func (u *ItodenwaUsecase) UseTelephoneCard() error {
	panic("unimplemented")
}

type PoolUserConn struct {
	PC     service.Websocket
	Mobile service.Websocket
}

type PoolUser struct {
	PC     entity.User
	Mobile entity.User
}

func NewItodenwaUsecase(roomRepository repository.IRoomRepository) IItodenwaUsecase {
	return &ItodenwaUsecase{
		roomRepository: roomRepository,
		Pool:           make(map[string]PoolUserConn),
		UserMap:        make(map[string]PoolUser),
		mu:             sync.RWMutex{},
	}
}

func (u *ItodenwaUsecase) CreateRoom(ctx context.Context, callerTel, calleeTel string) error {
	caller := u.GetUser(ctx, callerTel, false)
	if caller == nil {
		slog.Error("Caller not found in pool", "callerTel", callerTel)
		return errors.New("caller not found in pool")
	}
	callee := u.GetUser(ctx, calleeTel, false)
	if callee == nil {
		slog.Error("Callee not found in pool", "calleeTel", calleeTel)
		return errors.New("callee not found in pool")
	}
	uid := uuid.New(ctx)
	u.setRoomMap(ctx, calleeTel, uid)
	u.setRoomMap(ctx, callerTel, uid)
	r := entity.Room{
		ID:     uid,
		Caller: *caller,
		Callee: *callee,
	}
	err := u.roomRepository.Create(r)
	if err != nil {
		return err
	}
	callerPCConn := u.GetPool(ctx, callerTel, true)
	u.roomRepository.AddConnection(r.ID, caller, callerPCConn, true, true)
	callerMobileConn := u.GetPool(ctx, callerTel, false)
	u.roomRepository.AddConnection(r.ID, caller, callerMobileConn, true, false)
	calleeMobileConn := u.GetPool(ctx, calleeTel, false)
	u.roomRepository.AddConnection(r.ID, callee, calleeMobileConn, true, false)
	return nil
}

func (u *ItodenwaUsecase) AddPool(ctx context.Context, user entity.User, conn service.Websocket, isPC bool) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, exists := u.Pool[user.Tel]; exists {
		slog.Info("User already exists in pool, overwrite", "user", user.Tel)
	}
	if isPC {
		u.Pool[user.Tel] = PoolUserConn{
			PC:     conn,
			Mobile: u.Pool[user.Tel].Mobile,
		}
		u.UserMap[user.Tel] = PoolUser{
			PC:     user,
			Mobile: u.UserMap[user.Tel].Mobile,
		}
	} else {
		u.Pool[user.Tel] = PoolUserConn{
			PC:     u.Pool[user.Tel].PC,
			Mobile: conn,
		}
		u.UserMap[user.Tel] = PoolUser{
			PC:     u.UserMap[user.Tel].PC,
			Mobile: user,
		}
	}
	return nil
}

func (u *ItodenwaUsecase) RemovePool(ctx context.Context, user entity.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, exists := u.Pool[user.Tel]; !exists {
		slog.Info("User not found in pool", "user", user.Tel)
		return nil
	}

	delete(u.Pool, user.Tel)
	slog.Info("User removed from pool", "user", user.Tel)
	return nil
}

func (u *ItodenwaUsecase) GetPool(ctx context.Context, tel string, isPC bool) *service.Websocket {
	u.mu.RLock()
	defer u.mu.RUnlock()

	conn, exists := u.Pool[tel]
	if !exists {
		slog.Info("User not found in pool", "user", tel)
		return nil
	}

	slog.Info("User found in pool", "user", tel)
	if isPC {
		return &conn.PC
	}
	return &conn.Mobile
}
func (u *ItodenwaUsecase) GetUser(ctx context.Context, tel string, isPC bool) *entity.User {
	u.mu.RLock()
	defer u.mu.RUnlock()

	user, exists := u.UserMap[tel]
	if !exists {
		slog.Info("User not found in pool", "user", tel)
		return nil
	}

	slog.Info("User found in pool", "user", tel)
	if isPC {
		return &user.PC
	}
	return &user.Mobile
}

func (u *ItodenwaUsecase) Calling(ctx context.Context, user entity.User, conn service.Websocket) error {
	roomID, err := u.getRoomMap(ctx, user.Tel)
	if err != nil {
		slog.Error("Failed to get room ID for user", "user", user.Tel, "error", err)
		return err
	}
	binaryCh, jsonCh, errCh, err := u.roomRepository.GetChannels(roomID)
	if err != nil {
		slog.Error("Failed to get channels for user", "user", user.Tel, "error", err)
		return err
	}

	callerPCConn := u.GetPool(ctx, user.Tel, true)
	if callerPCConn == nil {
		slog.Error("PC connection not found for user", "user", user.Tel)
		return errors.New("PC connection not found")
	}
	callerMobileConn := u.GetPool(ctx, user.Tel, false)
	if callerMobileConn == nil {
		slog.Error("Mobile connection not found for user", "user", user.Tel)
		return errors.New("Mobile connection not found")
	}
	calleeMobileConn := u.GetPool(ctx, user.Tel, false)
	if calleeMobileConn == nil {
		slog.Error("Callee mobile connection not found for user", "user", user.Tel)
		return errors.New("Callee mobile connection not found")
	}
	go func() {
		for {
			callerPCConn.Read(ctx, jsonCh, binaryCh, errCh)
		}
	}()
	go func() {
		for {
			callerMobileConn.Read(ctx, jsonCh, binaryCh, errCh)
		}
	}()
	go func() {
		for {
			calleeMobileConn.Read(ctx, jsonCh, binaryCh, errCh)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-jsonCh:
				slog.Info("Received JSON message", "message", msg)
				if err := callerPCConn.Send(ctx, websocket.TextMessage, []byte(msg)); err != nil {
					slog.Error("Failed to send JSON message to PC", "error", err)
				}
				if err := callerMobileConn.Send(ctx, websocket.TextMessage, []byte(msg)); err != nil {
					slog.Error("Failed to send JSON message to Mobile", "error", err)
				}
				if err := calleeMobileConn.Send(ctx, websocket.TextMessage, []byte(msg)); err != nil {
					slog.Error("Failed to send JSON message to Callee Mobile", "error", err)
				}
			case binary := <-binaryCh:
				slog.Info("Received binary message", "message", binary)
				if err := callerPCConn.Send(ctx, websocket.BinaryMessage, binary); err != nil {
					slog.Error("Failed to send binary message to PC", "error", err)
				}
				if err := callerMobileConn.Send(ctx, websocket.BinaryMessage, binary); err != nil {
					slog.Error("Failed to send binary message to Mobile", "error", err)
				}
				if err := calleeMobileConn.Send(ctx, websocket.BinaryMessage, binary); err != nil {
					slog.Error("Failed to send binary message to Callee Mobile", "error", err)
				}
			case err := <-errCh:
				if err != nil {
					slog.Error("Error occurred in WebSocket connection", "error", err)
					if err := callerPCConn.Close(); err != nil {
						slog.Error("Failed to close PC connection", "error", err)
					}
					if err := callerMobileConn.Close(); err != nil {
						slog.Error("Failed to close Mobile connection", "error", err)
					}
					if err := calleeMobileConn.Close(); err != nil {
						slog.Error("Failed to close Callee Mobile connection", "error", err)
					}
					return
				}
			}
		}
	}()

	return nil
}

func (u *ItodenwaUsecase) Receive(ctx context.Context, conn service.Websocket) error {
	// Implementation for receiving data from the websocket connection
	return nil
}

func (u *ItodenwaUsecase) setRoomMap(ctx context.Context, tel string, roomID string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.RoomMap == nil {
		u.RoomMap = make(map[string]string)
	}
	u.RoomMap[tel] = roomID
	return nil
}
func (u *ItodenwaUsecase) getRoomMap(ctx context.Context, tel string) (string, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.RoomMap == nil {
		return "", errors.New("RoomMap is not initialized")
	}
	roomID, exists := u.RoomMap[tel]
	if !exists {
		return "", errors.New("Room not found")
	}
	return roomID, nil
}
