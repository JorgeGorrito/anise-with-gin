package command

type Listener interface {
	Listen(port uint16)
}
