package commands

const (
	SUCCESS_MESSAGE = iota
	FAIL_MESSAGE

	COMMAND_MESSAGE

	PRINT_MESSAGE
	PRINT_ERROR_MESSAGE
)

type Message struct {
	Type uint8  `json:"type"`
	Data []byte `json:"data"`
}
