package http

import (
	"botwhatsapp/internal/app/xodo"
	"botwhatsapp/internal/infra/ports"
	"botwhatsapp/internal/pkg/response"
	"botwhatsapp/util"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type XodoHttp struct {
	Log     ports.Logger
	gateway xodo.Gateway
	output  *response.Response
}

func NewXodoHttp(l ports.Logger, gateway xodo.Gateway) *XodoHttp {
	return &XodoHttp{
		Log:     l,
		output:  response.NewResponse("xodo", l),
		gateway: gateway,
	}
}

func (ctl *XodoHttp) rateHandler(w http.ResponseWriter, r *http.Request) {
	//input, err := util.ToStruct[](r.Body)
	input, err := util.ToStruct[any](r.Body)
	if err != nil {
		ctl.output.Error(input, err.Error(), w)
		return
	}

	output, err := ctl.gateway.Rate()
	if err != nil {
		ctl.output.Error(input, err.Error(), w)
		return
	}

	ctl.output.Success(input, output, w)
}

func (xd *XodoHttp) Handlers(r chi.Router) *XodoHttp {
	r.Route("/v1/xodo", func(r chi.Router) {
		r.Post("/rate", xd.rateHandler)
	})

	return xd
}
