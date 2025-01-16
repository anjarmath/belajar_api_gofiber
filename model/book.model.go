package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Judul     string     `json:"judul"`
	Deskripsi string     `json:"deskripsi"`
	Harga     int        `json:"harga"`
	CreatedAt time.Time  `gorm:"default:now();" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

var BooksCollection []Book

func InitBook() {
	BooksCollection = []Book{}
}
