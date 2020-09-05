package database

import "database/sql"

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) GetDB() *sql.DB {
	return s.DB
}

type Store interface {
	GetDB() *sql.DB
}

var _ Store = &SqlStore{}
