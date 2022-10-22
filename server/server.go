package server

import (
	"Api_With_AbdulHamid/domain/Phone"
	"Api_With_AbdulHamid/service"
)

type Server struct {
	service      service.Service
	phoneFactory Phone.Factory
}
