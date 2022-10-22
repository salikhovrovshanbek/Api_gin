package Phone

import "Api_With_AbdulHamid/pkg/id"

type Factory struct {
	idGenerator id.IGenerator
}

func NewFactory(ig id.IGenerator) Factory {
	return Factory{idGenerator: ig}
}

func (f Factory) NewPhone(name string) (Phone, error) {
	p := Phone{
		id:   f.idGenerator.GenerateUUID(),
		name: name,
	}

	if err := p.validate(); err != nil {
		return Phone{}, err
	}
	return p, nil
}
