package service

import "Api_With_AbdulHamid/domain/Phone"

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
