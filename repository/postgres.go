package repository

import (
	"Api_With_AbdulHamid/config"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(cfg config.PostgresConfig) (*Postgres, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}
