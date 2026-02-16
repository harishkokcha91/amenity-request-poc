package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func New(dsn string) (*DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (d *DB) Close() error {
	return d.DB.Close()
}
