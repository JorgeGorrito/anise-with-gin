package commands

type Registry interface {
	Register(name string, command Command)
}
