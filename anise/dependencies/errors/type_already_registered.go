package errors

type ErrTypeAlreadyRegistered struct {
	typeName string
}

func NewErrTypeAlreadyRegistered(typeName string) *ErrTypeAlreadyRegistered {
	return &ErrTypeAlreadyRegistered{
		typeName: typeName,
	}
}

func (e *ErrTypeAlreadyRegistered) Error() string {
	return "Type already registered: " + e.typeName
}
