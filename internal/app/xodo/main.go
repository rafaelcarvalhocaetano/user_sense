package xodo

import (
	"botwhatsapp/internal/app/xodo/usecase"
	"botwhatsapp/internal/interfaces/services"
	"botwhatsapp/internal/interfaces/webhooks/model"
)

type MainXodo struct{}

func New() *MainXodo {
	return &MainXodo{}
}

func (*MainXodo) Main(http services.WTAGateway, cc chan model.Channel) *Gateway {
	rate := usecase.NewRate(http, cc)
	message := usecase.NewMessage(http)
	mkt := usecase.NewMarketing(http)

	xodoGateway := Gateway{
		RateGateway:      rate,
		MarketingGateway: mkt,
		MessageGateway:   message,
	}

	return &xodoGateway
}
