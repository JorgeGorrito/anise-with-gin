package errors

type ErrTypeNotRegisterInDependencyProvider struct {
	typeName string
}

func NewErrTypeNotRegisterInDependencyProvider(typeName string) *ErrTypeNotRegisterInDependencyProvider {
	return &ErrTypeNotRegisterInDependencyProvider{
		typeName: typeName,
	}
}

func (e *ErrTypeNotRegisterInDependencyProvider) Error() string {
	return "Type not registered: " + e.typeName
}
