package xodo

import (
	"botwhatsapp/internal/app/xodo/usecase"
	"botwhatsapp/internal/interfaces/services"
)

type MainXodo struct{}

func New() *MainXodo {
	return &MainXodo{}
}

func (*MainXodo) Main(http services.WTAGateway) *Gateway {
	rate := usecase.NewRate(http)

	xodoGateway := Gateway{
		RateGateway: rate,
	}

	return &xodoGateway
}
