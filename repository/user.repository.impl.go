package repository

import (
	db "github.com/anjarmath/01_golang_api_sederhana/DB"
	"github.com/anjarmath/01_golang_api_sederhana/model"
)

type userRepositoryImpl struct{}

// AddUser implements UserRepository.
func (u userRepositoryImpl) AddUser(user *model.User) error {
	err := db.DB.Create(user).Error
	return err
}

// GetUserByID implements UserRepository.
func (u userRepositoryImpl) GetUserByUsername(username string) (*model.User, error) {
	user := new(model.User)
	if err := db.DB.Preload("FavoriteBook").Preload("Books").First(user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository() UserRepository {
	return userRepositoryImpl{}
}
