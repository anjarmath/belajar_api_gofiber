package repository

import "github.com/anjarmath/01_golang_api_sederhana/model"

type UserRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	AddUser(user *model.User) error
}
