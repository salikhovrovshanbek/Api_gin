package repository

import (
	"Api_With_AbdulHamid/config"
	"context"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func (p *Postgres) AddNewPhone(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) GetPhone(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) DeletePhone(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
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
