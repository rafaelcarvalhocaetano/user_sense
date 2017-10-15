package http

import (
	"botwhatsapp/internal/app/whatsapp"
	"botwhatsapp/internal/app/whatsapp/entities"
	wtaMessageDTO "botwhatsapp/internal/app/whatsapp/usecases/message/dto"
	"botwhatsapp/internal/app/whatsapp/usecases/templates/dto"
	"botwhatsapp/internal/infra/ports"
	"botwhatsapp/internal/pkg/response"
	"botwhatsapp/util"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type TemplatesHttp struct {
	log     ports.Logger
	output  *response.Response
	gateway whatsapp.Gateway
}

func NewTemplatesHttp(l ports.Logger, gateway whatsapp.Gateway) *TemplatesHttp {
	return &TemplatesHttp{
		log:     l,
		output:  response.NewResponse("template", l),
		gateway: gateway,
	}
}

func (tmp *TemplatesHttp) registerHandler(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[entities.TemplateModel](r.Body)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	resp, err := tmp.gateway.ExecuteRegister(input)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
	}

	tmp.output.Success(input, resp, w)
}

func (tmp *TemplatesHttp) dispatchTemplate(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[dto.Dispatch](r.Body)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	resp, err := tmp.gateway.ExecuteDispatch(input)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	tmp.output.Success(input, resp, w)
}

func (tmp *TemplatesHttp) getTemplateModelMeta(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("template_name")
	companyID := r.URL.Query().Get("company_id")
	input := make(map[string]string)
	if name == "" || companyID == "" {
		tmp.output.Error(input, "name or company_id is required", w)
		return
	}

	input["name"] = name
	input["company_id"] = companyID
	resp, err := tmp.gateway.GetTemplateMeta(input)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	tmp.output.Success(input, resp, w)
}

func (tmp *TemplatesHttp) sendTemplate(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[map[string]any](r.Body)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	resp, err := tmp.gateway.SendTemplate(input)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	tmp.output.Success(input, resp, w)
}

func (tmp *TemplatesHttp) simpleMessage(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[wtaMessageDTO.MessageData](r.Body)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	resp, err := tmp.gateway.SimpleMessage(input)
	if err != nil {
		tmp.output.Error(input, err.Error(), w)
		return
	}

	tmp.output.Success(input, resp, w)
}

func (tmp *TemplatesHttp) Handlers(r chi.Router) *TemplatesHttp {
	r.Route("/v1/templates", func(r chi.Router) {
		r.Post("/", tmp.registerHandler)
		r.Get("/data-model", tmp.getTemplateModelMeta)
		r.Post("/dispatch", tmp.dispatchTemplate)
		r.Post("/send", tmp.sendTemplate)
		r.Post("/simple-message", tmp.simpleMessage)
	})
	return tmp
}
