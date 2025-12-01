package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"` // "-" agar password tidak tampil di JSON response
	Role      string         `gorm:"type:varchar(20);default:'user'" json:"role"`
	Phone     string         `gorm:"type:varchar(20);default:null" json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Support Soft Delete
}

// TableName memaksa nama tabel menjadi 'users'
func (User) TableName() string {
	return "users"
}
