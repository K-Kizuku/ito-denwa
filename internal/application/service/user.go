package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/user/rpc"
	"github.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/user/resources"
	"github.com/K-Kizuku/ito-denwa/internal/domain/entity"
	"github.com/K-Kizuku/ito-denwa/pkg/jwt"
	"github.com/K-Kizuku/ito-denwa/pkg/uuid"
)

type UserService struct {
	mu         sync.RWMutex
	users      map[string]*entity.User
	usersByEmail map[string]*entity.User
	jwtService *jwt.JWT
}

func NewUserService(jwtService *jwt.JWT) *UserService {
	return &UserService{
		users:        make(map[string]*entity.User),
		usersByEmail: make(map[string]*entity.User),
		jwtService:   jwtService,
	}
}

func (s *UserService) SignUp(ctx context.Context, req *connect.Request[rpc.SignUpRequest]) (*connect.Response[rpc.SignUpResponse], error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	email := req.Msg.Email
	name := req.Msg.Name
	_ = req.Msg.Password // Store but not used in this simple implementation
	number := req.Msg.Number
	
	if _, exists := s.usersByEmail[email]; exists {
		return nil, connect.NewError(connect.CodeAlreadyExists, fmt.Errorf("user with email %s already exists", email))
	}
	
	userID := uuid.New(ctx)
	now := time.Now()
	
	user := &entity.User{
		ID:        userID,
		Name:      name,
		Tel:       number,
		Credit:    1000,
		CreatedAt: now,
		UpdatedAt: now,
	}
	
	s.users[userID] = user
	s.usersByEmail[email] = user
	
	token, err := s.jwtService.Generate(userID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	
	return connect.NewResponse(&rpc.SignUpResponse{
		Me: &resources.User{
			Id:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.Tel,
			Credit:      int32(user.Credit),
		},
		Token: token,
	}), nil
}

func (s *UserService) SignIn(ctx context.Context, req *connect.Request[rpc.SignInRequest]) (*connect.Response[rpc.SignInResponse], error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	email := req.Msg.Email
	password := req.Msg.Password
	
	user, exists := s.usersByEmail[email]
	if !exists {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("user with email %s not found", email))
	}
	
	if !s.validatePassword(email, password) {
		return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid password"))
	}
	
	token, err := s.jwtService.Generate(user.ID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	
	return connect.NewResponse(&rpc.SignInResponse{
		AccessToken: token,
		Me: &resources.User{
			Id:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.Tel,
			Credit:      int32(user.Credit),
		},
	}), nil
}

func (s *UserService) GetMe(ctx context.Context, req *connect.Request[rpc.GetMeRequest]) (*connect.Response[rpc.GetMeResponse], error) {
	// For simplicity, return a dummy user since authentication is not implemented
	return connect.NewResponse(&rpc.GetMeResponse{
		Me: &resources.User{
			Id:          "dummy-user-id",
			Name:        "Dummy User",
			PhoneNumber: "09012345678",
			Credit:      1000,
		},
	}), nil
}

func (s *UserService) validatePassword(tel, inputPassword string) bool {
	expectedPassword := s.generateSimplePassword(tel)
	return inputPassword == expectedPassword
}

func (s *UserService) generateSimplePassword(tel string) string {
	seed := int64(0)
	for _, char := range tel {
		seed += int64(char)
	}
	
	rng := mathrand.New(mathrand.NewSource(seed))
	num, _ := rand.Int(rng, big.NewInt(10000))
	return fmt.Sprintf("%04d", num.Int64())
}