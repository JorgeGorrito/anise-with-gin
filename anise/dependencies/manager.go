package dependencies

type Manager interface {
	SetResolver(dependenciesResolver Resolver)
	GetResolver() Resolver
	RegisterDependencies(dependencyBinder Binder) error
}
