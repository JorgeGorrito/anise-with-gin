package commands

import "fmt"

type Factory struct {
	registry map[string]Command
}

func NewFactory() *Factory {
	return &Factory{
		registry: make(map[string]Command),
	}
}

func (f *Factory) Register(name string, command Command) {
	if _, ok := f.registry[name]; ok {
		panic(fmt.Sprintf("Command %s already registered", name))
	}
	f.registry[name] = command
}

func (f *Factory) GetByName(name string) Command {
	if command, ok := f.registry[name]; ok {
		return command
	}
	return nil
}
