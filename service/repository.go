package service

import "context"

type Repository interface {
	AddNewPhone(ctx context.Context) error
	GetPhone(ctx context.Context) error
	DeletePhone(ctx context.Context) error
}
