package command

type Registry interface {
	Register(name string, construct NewCommand)
}
