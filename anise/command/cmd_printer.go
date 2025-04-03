package command

type CMDPrinter interface {
	Print(a ...any)
	Println(a ...any)
	Printf(format string, a ...any)
}
