package command

import "fmt"

type NewCommand func(params map[string]string, printer CMDPrinter) Command

type Factory interface {
	Register(name string, construct NewCommand)
	GetByName(name string) NewCommand
}

type factory struct {
	registry map[string]NewCommand
}

func NewFactory() *factory {
	return &factory{
		registry: make(map[string]NewCommand),
	}
}

func (f *factory) Register(name string, construct NewCommand) {
	if _, ok := f.registry[name]; ok {
		panic(fmt.Sprintf("Command %s already registered", name))
	}
	f.registry[name] = construct
}

func (f *factory) GetByName(name string) NewCommand {
	if construct, ok := f.registry[name]; ok {
		return construct
	}
	return nil
}
