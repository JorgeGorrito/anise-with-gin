package errors

import "errors"

var (
	ErrDependenciesManagerIsNil error = errors.New("dependencies manager can't be nil")
	ErrEngineIsNil              error = errors.New("gin engine can't be nil")
	ErrRoutesManagerIsNil       error = errors.New("routing manager can't be nil")
	ErrConfigManagerIsNil       error = errors.New("config manager can't be nil")
)
