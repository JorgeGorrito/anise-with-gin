package commands

import (
	"encoding/json"
	"fmt"
	"net"
)

type CMD struct {
	conn net.Conn
}

func NewCMD(conn net.Conn) *CMD {
	return &CMD{
		conn: conn,
	}
}

func (c *CMD) send(message Message) {
	messageToSend, err := json.Marshal(&message)
	if err != nil {
		panic(err)
	}
	_, err = c.conn.Write([]byte(messageToSend))
	if err != nil {
		panic(err)
	}
}

func (c *CMD) Print(a ...any) {
	var outCmd = fmt.Sprint(a...)
	c.send(Message{
		Type: PRINT_MESSAGE,
		Data: []byte(outCmd),
	})
}

func (c *CMD) Println(a ...any) {
	var outCmd = fmt.Sprintln(a...)
	c.send(Message{
		Type: PRINT_MESSAGE,
		Data: []byte(outCmd),
	})
}

func (c *CMD) Printf(format string, a ...any) {
	var outCmd = fmt.Sprintf(format, a...)
	c.send(Message{
		Type: PRINT_MESSAGE,
		Data: []byte(outCmd),
	})
}
