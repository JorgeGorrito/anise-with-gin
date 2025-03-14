package dependencies

import (
	"reflect"

	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies/errors"
	"github.com/JorgeGorrito/anise-with-gin/anise/dependencies/types"
)

func GetAbstractType[T any]() types.Abstract {
	var abstractType reflect.Type = reflect.TypeOf((*T)(nil)).Elem()
	if abstractType.Kind() != reflect.Interface {
		panic(errors.ErrAbstractMustBeInterface)
	}
	return abstractType
}

func Inject[T any](dependencyResolver Resolver) T {
	return dependencyResolver.Resolve(GetAbstractType[T]()).(T)
}
