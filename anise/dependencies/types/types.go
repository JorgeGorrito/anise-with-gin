package types

import "reflect"

type Abstract reflect.Type
type Concrete any
type GetConcreteFunc func() any
