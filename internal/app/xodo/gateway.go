package xodo

type RateGateway interface {
	Rate() (any, error)
}

type Gateway struct {
	RateGateway
}
