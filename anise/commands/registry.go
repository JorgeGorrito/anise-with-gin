package commands

type Registry interface {
	Register(name string, construct NewCommand)
}
