package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Title     string
	Caption   string
	PhotoURL  string
	UserID    uint
	User      User // Relasi dengan model User
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Sesuaikan dengan atribut yang diperlukan sesuai deskripsi Anda.

type PhotoDTO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
