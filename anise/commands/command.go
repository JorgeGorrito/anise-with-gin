package commands

type Command interface {
	Execute(params map[string]string) error
}
