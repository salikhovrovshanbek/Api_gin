package server

import (
	"Api_With_AbdulHamid/domain/Phone"
	"Api_With_AbdulHamid/service"
)

type Server struct {
	service      service.Service
	phoneFactory Phone.Factory
}

func New(service2 service.Service, factory Phone.Factory) Server {
	return Server{
		service:      service2,
		phoneFactory: factory,
	}
}
