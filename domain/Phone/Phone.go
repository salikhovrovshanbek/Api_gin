package Phone

import (
	"errors"
	"github.com/google/uuid"
)

type Phone struct {
	id   uuid.UUID
	name string
}

func (p Phone) Id() uuid.UUID {
	return p.id
}
func (p Phone) Name() string {
	return p.name
}

func (p Phone) validate() error {
	if p.name == "" {
		return errors.New("Phone name is empty")
	}
	return nil
}
