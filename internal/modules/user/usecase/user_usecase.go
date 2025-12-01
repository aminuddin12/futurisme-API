package usecase

import (
	"context"
	"futurisme-api/internal/modules/user/entity"
	"futurisme-api/internal/modules/user/repository"
)

// UserUseCase Interface
type UserUseCase interface {
	GetProfile(ctx context.Context, userID uint) (*entity.User, error)
}

// Implementasi
type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) GetProfile(ctx context.Context, userID uint) (*entity.User, error) {
	// Memanggil repository untuk cari user by ID
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
