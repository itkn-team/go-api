package teststore

import (
	"github.com/shaolinjehzu/goAPI/internal/app/model"
	"github.com/shaolinjehzu/goAPI/internal/app/store"
)

// Store struct
type Store struct {
	UserRepository *UserRepository
}

// New Store
func New() *Store {
	return &Store{
	}
}


//User repository
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}

	return s.UserRepository


}
