package services

type SendMessagesHttp interface {
	Send(p string, data any) (*string, error)
}

type WAGateway struct {
	SendMessagesHttp
}
