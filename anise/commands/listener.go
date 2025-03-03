package commands

type Listener interface {
	Listen(port uint16)
}
