package main

import (
	"os"

	"github.com/JorgeGorrito/anise-with-gin/anise/commands"
)

func main() {
	if len(os.Args) > 1 {
		commandSender := commands.NewDefaultSender(os.Args)
		commandSender.Send()
	}
}
