package main

import (
	"os"

	"github.com/JorgeGorrito/anise-with-gin/anise/command"
)

func main() {
	if len(os.Args) > 1 {
		commandSender := command.NewDefaultSender(os.Args)
		commandSender.Send()
	}
}
