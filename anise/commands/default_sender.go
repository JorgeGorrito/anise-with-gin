package commands

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

type DefaultSender struct {
	args []string
}

func NewDefaultSender(args []string) *DefaultSender {
	return &DefaultSender{
		args: args,
	}
}

func (s *DefaultSender) getCommandRequest() CommandRequest {
	return CommandRequest{
		Name: s.args[1],
		Params: func(args []string) map[string]string {
			var params map[string]string = make(map[string]string)
			for _, param := range args {
				paramsFormatted := strings.Split(param[2:], "=")
				params[paramsFormatted[0]] = paramsFormatted[1]
			}

			return params
		}(s.args[2:]),
	}
}

func (s *DefaultSender) sendMessage(conn net.Conn, message *Message) {
	messageToSend, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte(messageToSend))
	if err != nil {
		panic(err)
	}
}

func (s *DefaultSender) Send() {
	var buffer []byte = make([]byte, 1024)
	var commandRequest CommandRequest = s.getCommandRequest()

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", DEFAULT_COMMAND_LISTENER_PORT))
	if err != nil {
		panic(fmt.Sprintf("[Anise-Commands] Error: It was not possible to establish a connection with the Anise application. \n%s\n", err.Error()))
	}
	defer conn.Close()

	dataToSend, err := json.Marshal(commandRequest)
	if err != nil {
		panic(fmt.Sprintf("[Anise-Commands] Error: It was not possible to decode the received message. \n%s\n", err.Error()))
	}

	s.sendMessage(conn, &Message{
		Type: COMMAND_MESSAGE,
		Data: dataToSend,
	})

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("[Anise-Commands] Error: %s\n", err)
			break
		}
		var message Message
		if err := json.Unmarshal(buffer[:n], &message); err != nil {
			fmt.Printf("[Anise-Commands] Error: %s\n", err)
			break
		}

		switch message.Type {
		case SUCCESS_MESSAGE:
			fmt.Printf("[Anise-Commands] Success: %s\n", commandRequest.Name)
			return
		case FAIL_MESSAGE:
			fmt.Printf("[Anise-Commands] Fail %s: %s\n", commandRequest.Name, string(message.Data))
			return
		case PRINT_MESSAGE:
			fmt.Printf("[Anise-Commands] %s: %s\n", commandRequest.Name, string(message.Data))
		default:
			continue
		}
	}
}
