package commands

type Manager interface {
	RegisterCommands(registry Registry) error
}
