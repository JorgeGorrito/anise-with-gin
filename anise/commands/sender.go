package commands

type Sender interface {
	Send(request CommandRequest)
}
