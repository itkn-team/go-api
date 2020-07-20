package store

import "github.com/shaolinjehzu/goAPI/internal/app/model"

//UserRepository
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
