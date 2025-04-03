package command

type Retriever interface {
	GetByName(name string) NewCommand
}
