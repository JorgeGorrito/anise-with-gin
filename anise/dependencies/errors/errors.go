package errors

import "errors"

var (
	ErrConcreteImplementIsntPoint error = errors.New("concrete implement isn't a pointer")
	ErrAbstractMustBeInterface    error = errors.New("abstract must be an interface")
)
