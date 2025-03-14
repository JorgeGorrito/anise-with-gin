package errors

type ErrTypeNotRegister struct {
	typeName string
}

func NewErrTypeNotRegister(typeName string) *ErrTypeNotRegister {
	return &ErrTypeNotRegister{
		typeName: typeName,
	}
}

func (e *ErrTypeNotRegister) Error() string {
	return "Type not registered: " + e.typeName
}
