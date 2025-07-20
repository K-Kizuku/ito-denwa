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
)

type IItodenwaUsecase interface {
	Receive(ctx context.Context, conn service.Websocket) error
	Send(ctx context.Context, conn service.Websocket) error
	Processing() error
	Calling() error
	AddPool(ctx context.Context, user entity.User, conn service.Websocket) error
	RemovePool(ctx context.Context, user entity.User) error
	GetPool(ctx context.Context, tel string) *service.Websocket
	CreateRoom(ctx context.Context) error

	UseTelephoneCard() error
}

type ItodenwaUsecase struct {
	roomRepository repository.IRoomRepository
	Pool           map[string]service.Websocket
	mu             sync.RWMutex
}

func NewItodenwaUsecase(roomRepository repository.IRoomRepository) IItodenwaUsecase {
	return &ItodenwaUsecase{
		roomRepository: roomRepository,
		Pool:           make(map[string]service.Websocket),
		mu:             sync.RWMutex{},
	}
}

func (u *ItodenwaUsecase) CreateRoom(ctx context.Context) error {
	r := entity.Room{
		ID: uuid.New(ctx),
	}
	if r.ID == "" {
		slog.Error("Room ID is not provided in context")
		return errors.New("room ID is required")
	}
	err := u.roomRepository.Create(r)
	if err != nil {
		return err
	}
	return nil
}

func (u *ItodenwaUsecase) AddPool(ctx context.Context, user entity.User, conn service.Websocket) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if _, exists := u.Pool[user.Tel]; exists {
		slog.Info("User already exists in pool, overwrite", "user", user.Tel)
	}

	u.Pool[user.Tel] = conn
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

func (u *ItodenwaUsecase) GetPool(ctx context.Context, tel string) *service.Websocket {
	u.mu.RLock()
	defer u.mu.RUnlock()

	conn, exists := u.Pool[tel]
	if !exists {
		slog.Info("User not found in pool", "user", tel)
		return nil
	}

	slog.Info("User found in pool", "user", tel)
	return &conn
}

func (u *ItodenwaUsecase) Calling(ctx context.Context, user entity.User, conn service.Websocket) error {
	if err := u.AddPool(ctx, user, conn); err != nil {
		slog.Error("Failed to add user to pool", "user", user.Tel, "error", err)
		return err
	}
	slog.Info("User added to pool", "user", user.Tel)

	return nil
}

func (u *ItodenwaUsecase) Receive(ctx context.Context, conn service.Websocket) error {
	// Implementation for receiving data from the websocket connection
	return nil
}
