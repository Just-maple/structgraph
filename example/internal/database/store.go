package database

import (
	"database/sql"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) GetDB() *sql.DB {
	return s.DB
}

type Store interface {
	GetDB() *sql.DB
}

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}

type Redis struct {
}

func (r Redis) Get(key string) (interface{}, error) {
	panic("implement me")
}

func (r Redis) Set(key string, value interface{}) error {
	panic("implement me")
}

var _ Cache = &Redis{}
var _ Store = &SqlStore{}
