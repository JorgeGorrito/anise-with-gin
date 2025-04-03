package command

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

func (l *CommandListener) sendMessage(connAccepted net.Conn, message *Message) {
	messageToSend, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("[Anise-Commands] Error: %s\n", err)
	}
	_, err = connAccepted.Write(messageToSend)
	if err != nil {
		fmt.Printf("[Anise-Commands] Error: %s\n", err)
	}
}

func (l *CommandListener) executeCommand(connAccepted net.Conn, commandToExecute Command) {
	var err error = commandToExecute.Execute()
	if err != nil {
		l.sendMessage(connAccepted, &Message{
			Type: FAIL_MESSAGE,
			Data: []byte(err.Error()),
		})
		return
	}
	l.sendMessage(connAccepted, &Message{
		Type: SUCCESS_MESSAGE,
		Data: nil,
	})
}

func (l *CommandListener) getMessageReceive(connAccepted net.Conn, messageReceived *Message) error {
	var buffer []byte = make([]byte, 1024)
	n, err := connAccepted.Read(buffer)
	if err != nil {
		l.sendMessage(connAccepted, &Message{
			Type: FAIL_MESSAGE,
			Data: []byte(fmt.Sprintf("<Error> %s\n", err)),
		})
		return err
	}
	if err := json.Unmarshal(buffer[:n], &messageReceived); err != nil {
		l.sendMessage(connAccepted, &Message{
			Type: FAIL_MESSAGE,
			Data: []byte(fmt.Sprintf("<Error> %s\n", err)),
		})
		return err
	}
	return nil
}

func (l *CommandListener) handleConnection(connAccepted net.Conn) {
	defer connAccepted.Close()
	var messageReceived Message
	var commandRequest CommandRequest

	if err := l.getMessageReceive(connAccepted, &messageReceived); err != nil {
		return
	}
	if messageReceived.Type != COMMAND_MESSAGE {
		l.sendMessage(connAccepted, &Message{
			Type: FAIL_MESSAGE,
			Data: []byte("<Error> Invalid message type"),
		})
		return
	}
	if err := json.Unmarshal(messageReceived.Data, &commandRequest); err != nil {
		l.sendMessage(connAccepted, &Message{
			Type: FAIL_MESSAGE,
			Data: []byte(fmt.Sprintf("<Error> %s", err)),
		})
		return
	}
	var InitCommand NewCommand = l.commandRetriever.GetByName(commandRequest.Name)
	if InitCommand == nil {
		l.sendMessage(connAccepted, &Message{
			Type: FAIL_MESSAGE,
			Data: []byte(fmt.Sprintf("<Error> Command %s not found", commandRequest.Name)),
		})
		return
	}

	l.executeCommand(connAccepted, InitCommand(commandRequest.Params, NewCMD(connAccepted)))
}

func (l *CommandListener) Listen(port uint16) {
	fmt.Printf("[Anise-Commands] Listening commands on port :%d\n", port)
	if listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("[Anise-Commands] Error: %s\n", err)
	} else {
		for {
			if conn, err := listener.Accept(); err != nil {
				fmt.Printf("[Anise-Commands] Error: %s\n", err)
			} else {
				go l.handleConnection(conn)
			}
		}
	}
}
