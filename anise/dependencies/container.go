package dependencies

import (
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies/errors"
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies/types"
)

type Container struct {
	registry map[types.Abstract]types.GetConcreteFunc
}

func NewContainer() *Container {
	return &Container{
		registry: make(map[types.Abstract]types.GetConcreteFunc),
	}
}

func (c *Container) Bind(abstract types.Abstract, getConcreteFunc types.GetConcreteFunc) {
	if _, ok := c.registry[abstract]; ok {
		panic(errors.NewErrTypeAlreadyRegistered(abstract.String()))
	}
	c.registry[abstract] = getConcreteFunc
}

func (c *Container) Resolve(abstract types.Abstract) types.Concrete {
	getConcreteFunc, ok := c.registry[abstract]
	if !ok {
		panic(errors.NewErrTypeNotRegister(abstract.String()))
	}
	return getConcreteFunc()
}
