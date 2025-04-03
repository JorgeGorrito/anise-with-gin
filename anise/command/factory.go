package command

import "fmt"

type NewCommand func(params map[string]string, printer CMDPrinter) Command

type Factory struct {
	registry map[string]NewCommand
}

func NewFactory() *Factory {
	return &Factory{
		registry: make(map[string]NewCommand),
	}
}

func (f *Factory) Register(name string, construct NewCommand) {
	if _, ok := f.registry[name]; ok {
		panic(fmt.Sprintf("Command %s already registered", name))
	}
	f.registry[name] = construct
}

func (f *Factory) GetByName(name string) NewCommand {
	if construct, ok := f.registry[name]; ok {
		return construct
	}
	return nil
}
