package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Username  string  `gorm:"not null"`
	Email     string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Photos    []Photo // Relasi dengan model Photo
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Sesuaikan dengan atribut yang diperlukan sesuai deskripsi Anda.

type UserDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
