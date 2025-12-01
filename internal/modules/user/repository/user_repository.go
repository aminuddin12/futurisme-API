package repository

import (
	"context"
	"errors"
	"futurisme-api/internal/modules/user/entity"

	"gorm.io/gorm"
)

// UserRepository Interface mendefinisikan kontrak fungsi apa saja yang tersedia
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id uint) (*entity.User, error)
}

// userRepository adalah struct implementasi dari interface di atas
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository adalah constructor untuk membuat instance repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Implementasi fungsi Create (Register)
func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// Implementasi fungsi FindByEmail (Login)
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	// Mencari user berdasarkan email
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak error, tapi data kosong
		}
		return nil, err
	}
	return &user, nil
}

// Implementasi fungsi FindByID (Profile/Detail)
func (r *userRepository) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
