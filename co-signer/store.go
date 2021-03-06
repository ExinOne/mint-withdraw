package main

import (
	"context"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

// ErrNotFound not found
var ErrNotFound = errors.New("value not found")

// Store store
type Store struct {
	factory *cache.Cache
}

// NewStore new store
func NewStore() *Store {
	return &Store{
		cache.New(time.Hour, time.Hour),
	}
}

// ReadProperty read property
func (s Store) ReadProperty(ctx context.Context, key string) (string, error) {
	if s, f := s.factory.Get(key); f {
		return s.(string), nil
	}
	return "", ErrNotFound
}

// WriteProperty write property
func (s Store) WriteProperty(ctx context.Context, key, value string) error {
	s.factory.Set(key, value, time.Hour)
	return nil
}
