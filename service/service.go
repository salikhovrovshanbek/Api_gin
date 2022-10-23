package service

import (
	"Api_With_AbdulHamid/domain/Phone"
	"context"
)

type Service struct {
	repo         Repository
	PhoneFactory Phone.Factory
}

func New(repo Repository, phf Phone.Factory) Service {
	return Service{
		repo:         repo,
		PhoneFactory: phf,
	}
}

func (s Service) AddNewPhone(ctx context.Context) error {
	return nil
}

func (s Service) GetPhone(ctx context.Context) error {
	return nil
}

func (s Service) DeletePhone(ctx context.Context) error {
	return nil
}
