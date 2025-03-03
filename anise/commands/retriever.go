package commands

type Retriever interface {
	GetByName(name string) Command
}
