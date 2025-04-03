package command

type Sender interface {
	Send(request CommandRequest)
}
