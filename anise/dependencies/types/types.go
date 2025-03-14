package types

import "reflect"

type Abstract reflect.Type
type Concrete interface{}
type GetConcreteFunc func() any
