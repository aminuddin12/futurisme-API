package usecase

import (
	"context"
	"errors"
	"futurisme-api/config"
	"futurisme-api/internal/modules/user/entity"
	"futurisme-api/internal/modules/user/repository"
	jwtUtil "futurisme-api/pkg/utils/jwt"
	"futurisme-api/pkg/utils/security"
)

// DTOs (Data Transfer Objects)
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  *entity.User `json:"user"`
}

// Interface AuthUseCase
type AuthUseCase interface {
	Register(ctx context.Context, req *RegisterRequest) (*entity.User, error)
	Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

// Implementasi AuthUseCase
type authUseCase struct {
	userRepo repository.UserRepository
	cfg      *config.Config
}

func NewAuthUseCase(userRepo repository.UserRepository, cfg *config.Config) AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

// Logika Register
func (uc *authUseCase) Register(ctx context.Context, req *RegisterRequest) (*entity.User, error) {
	// 1. Cek User Existing
	existingUser, _ := uc.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// 2. Hash Password
	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 3. Simpan User Baru
	newUser := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone:    req.Phone,
		Role:     "user",
	}

	if err := uc.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// Logika Login
func (uc *authUseCase) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// 1. Cari User berdasarkan Email
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// 2. Cek Password
	if !security.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// 3. Generate JWT Token
	token, err := jwtUtil.GenerateToken(user.ID, user.Role, uc.cfg.Security.JWTSecret, 24) // Assuming 24 hours expiry for now
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User:  user,
	}, nil
}
