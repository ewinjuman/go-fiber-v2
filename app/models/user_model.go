package models

import (
	"gorm.io/gorm"
	"time"
)

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	UserRole string `json:"userRole" validate:"required,lte=25"`
}

type SignUpResponse struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Status            int    `json:"status"`
	UserRole          string `json:"userRole"`
	OldID             int    `json:"oldId"`
	MobilePhoneNumber string `json:"mobilePhoneNumber"`
}

// User struct to describe User object.
type User struct {
	ID           int            `gorm:"column:id" json:"id" validate:"required,numeric"`
	Email        string         `gorm:"column:email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string         `gorm:"column:password_hash" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus   int            `gorm:"column:user_status" json:"user_status" validate:"required,len=1"`
	UserRole     string         `gorm:"column:user_role" json:"user_role" validate:"required,lte=25"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    *time.Time     `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
