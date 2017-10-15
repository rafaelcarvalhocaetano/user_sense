package http

import (
	"botwhatsapp/internal/app/whatsapp"
	wtaMessage "botwhatsapp/internal/app/whatsapp/usecases/message/dto"
	"botwhatsapp/internal/infra/ports"
	"botwhatsapp/internal/pkg/response"
	"botwhatsapp/util"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type InteractiveMessagesHttp struct {
	log     ports.Logger
	output  *response.Response
	gateway whatsapp.Gateway
}

func NewInteractiveMessagesHttp(
	l ports.Logger,
	gateway whatsapp.Gateway,
) *InteractiveMessagesHttp {
	return &InteractiveMessagesHttp{
		log:     l,
		output:  response.NewResponse("interactive_messages", l),
		gateway: gateway,
	}
}

func (ctl *InteractiveMessagesHttp) sendSimpleMessage(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[wtaMessage.MessageData](r.Body)
	if err != nil {
		ctl.output.Error(input, err.Error(), w)
		return
	}

	resp, err := ctl.gateway.SimpleMessage(input)
	if err != nil {
		ctl.output.Error(input, err.Error(), w)
		return
	}

	ctl.output.Success(input, resp, w)
}

func (ctl *InteractiveMessagesHttp) Handlers(r chi.Router) *InteractiveMessagesHttp {
	r.Route("/v1/message", func(r chi.Router) {
		r.Post("/simple", ctl.sendSimpleMessage)
	})
	return ctl
}
