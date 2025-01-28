package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `gorm:"default:now();" json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u *User) HashPassword() error {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}
	u.Password = string(hashedPw)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
