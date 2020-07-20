package sqlstore

import (
	"database/sql"
	"github.com/shaolinjehzu/goAPI/internal/app/store"
)

// Store struct
type Store struct {
	db *sql.DB
	UserRepository *UserRepository
}

// New Store
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}


//User repository
func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository


}