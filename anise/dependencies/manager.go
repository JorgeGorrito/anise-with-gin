package dependencies

type Manager interface {
	RegisterDependencies(dependencyBinder Binder) error
}
