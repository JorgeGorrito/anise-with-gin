package commands

type CommandRequest struct {
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
}
