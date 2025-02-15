package dependencies

import (
	"reflect"

	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies/errors"
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies/types"
)

type Container struct {
	registry map[types.Abstract]types.Concrete
}

func NewContainer() *Container {
	return &Container{
		registry: make(map[types.Abstract]types.Concrete),
	}
}

func (c *Container) Bind(abstract types.Abstract, concrete types.Concrete) {
	concreteType := reflect.TypeOf(concrete)
	if concreteType.Kind() != reflect.Ptr {
		panic(errors.ErrConcreteImplementIsntPoint)
	}

	c.registry[abstract] = concrete
}

func (c *Container) Resolve(abstract types.Abstract) types.Concrete {
	concrete, ok := c.registry[abstract]
	if !ok {
		panic(errors.NewErrTypeNotRegisterInDependencyProvider(abstract.String()))
	}
	return concrete
}
