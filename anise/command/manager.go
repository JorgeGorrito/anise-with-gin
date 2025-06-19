package command

type Manager interface {
	RegisterCommands(registry Registry) error
	GetRetrieverCommand() Retriever
	GetRegistryCommand() Registry
}
