package xodo

import "botwhatsapp/internal/app/xodo/dto"

type RateGateway interface {
	Rate(input dto.InputRate) (*string, error)
}

type MarketingGateway interface {
	Mkt(input dto.InputRate) (*string, error)
}

type MessageGateway interface {
	SendMessage(data *dto.InputMessage) (*dto.OutputMessage, error)
}

type Gateway struct {
	RateGateway
	MarketingGateway
	MessageGateway
}
