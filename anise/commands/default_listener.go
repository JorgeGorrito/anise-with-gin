package commands

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	DEFAULT_COMMAND_LISTENER_PORT = 7114
)

type CommandListener struct {
	commandRetriever Retriever
}

func NewDefaultListener(commandRetriever Retriever) *CommandListener {
	return &CommandListener{
		commandRetriever: commandRetriever,
	}
}

func (l *CommandListener) handleCommandReceive(connAccepted net.Conn) {
	defer connAccepted.Close()
	buffer := make([]byte, 1024)
	n, err := connAccepted.Read(buffer)
	if err != nil {
		fmt.Printf("Anise@Commands> Error: %s\n", err)
	}
	var commandRequest CommandRequest
	if err := json.Unmarshal(buffer[:n], &commandRequest); err != nil {
		fmt.Printf("Anise@Commands> Error: %s\n", err)
	}
	commandToExecute := l.commandRetriever.GetByName(commandRequest.Name)
	if commandToExecute == nil {
		fmt.Printf("Anise@Commands> Error: Command %s not found\n", commandRequest.Name)
		return
	}
	go commandToExecute.Execute(commandRequest.Params)
}

func (l *CommandListener) Listen(port uint16) {
	fmt.Printf("Anise@Commands> Listening commands on port %d\n", port)
	if listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("Anise@Commands> Error: %s\n", err)
	} else {
		for {
			if conn, err := listener.Accept(); err != nil {
				fmt.Printf("Anise@Commands> Error: %s\n", err)
			} else {
				go l.handleCommandReceive(conn)
			}
		}
	}
}
