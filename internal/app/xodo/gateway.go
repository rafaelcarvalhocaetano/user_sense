package xodo

import "botwhatsapp/internal/app/xodo/dto"

type RateGateway interface {
	Rate(input dto.Input) (any, error)
}

type Gateway struct {
	RateGateway
}
