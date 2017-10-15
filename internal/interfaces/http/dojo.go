package http

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"botwhatsapp/internal/app/whatsapp/usecases/dojo"
	"botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"
	"botwhatsapp/internal/infra/ports"
	"botwhatsapp/internal/pkg/response"
	"botwhatsapp/util"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type DojoHttp struct {
	log     ports.Logger
	output  *response.Response
	gateway dojo.Gateway
}

func NewDojoHttp(l ports.Logger, gateway dojo.Gateway) *DojoHttp {
	return &DojoHttp{
		log:     l,
		output:  response.NewResponse("dojo", l),
		gateway: gateway,
	}
}

func (dj *DojoHttp) createMsgInteractive(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[dto.InputDojoCreateMsgInteractiveDTO](r.Body)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	resp, err := dj.gateway.DojoCreateMsgInteractive(input)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
	}

	dj.output.Success(input, resp, w)
}

func (dj *DojoHttp) interactiveMsgByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		dj.output.Error(nil, "parametro inválido", w)
		return
	}

	resp, err := dj.gateway.GetModelInteractive(&id, nil)
	if err != nil {
		dj.output.Error(resp, err.Error(), w)

	}

	dj.output.Success(id, resp, w)
}

func (dj *DojoHttp) interactiveMsgByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		dj.output.Error(nil, "parametro inválido", w)
		return
	}

	resp, err := dj.gateway.GetModelInteractive(nil, &name)
	if err != nil {
		dj.output.Error(nil, err.Error(), w)
	}

	dj.output.Success(name, resp, w)
}

func (dj *DojoHttp) modelTemplateByWrite(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("template_name")
	if name == "" {
		dj.output.Error(name, "name or company_id is required", w)
		return
	}

	resp, err := dj.gateway.GetModelTemplateByWrite(name)
	if err != nil {
		dj.output.Error(name, err.Error(), w)
		return
	}

	dj.output.Success(name, resp, w)
}

func (dj *DojoHttp) registerTempalteMeta(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[entities.TemplateModel](r.Body)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	resp, err := dj.gateway.RegisterTemplateMeta(input)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	dj.output.Success(input, resp, w)

}

func (dj *DojoHttp) sendOneTemplate(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[map[string]any](r.Body)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	resp, err := dj.gateway.SendOneTemplate(*input)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
	}

	dj.output.Success(input, resp, w)
}

func (dj *DojoHttp) dispatchMessage(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[dto.DojoDispatchParams](r.Body)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	resp, err := dj.gateway.Dispatch(input)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	dj.output.Success(input, resp, w)
}

func (dj *DojoHttp) sendSimpleMessageHandler(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[dojo.InputSendSimpleMessage](r.Body)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	resp, err := dj.gateway.SendSimpleMessage(input)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	dj.output.Success(input, resp, w)
}

func (dj *DojoHttp) modelRegisterTemplate(w http.ResponseWriter, r *http.Request) {
	resp := dj.gateway.ModelRegisterTemplate()

	dj.output.Success(nil, resp, w)
}

func (dj *DojoHttp) createFlowHandler(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[*dto.InputCreateDojoFlow](r.Body)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	resp, err := dj.gateway.DojoCreateFlow(*input)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	dj.output.Success(input, resp, w)
}

func (dj *DojoHttp) getFlowHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		dj.output.Error(nil, "name is required params", w)
		return
	}

	metaID := r.URL.Query().Get("meta_id")
	if metaID == "" {
		dj.output.Error(nil, "meta_id is required", w)
		return
	}

	resp, err := dj.gateway.DojoGetFlow(metaID, name)
	if err != nil {
		dj.output.Error(name, err.Error(), w)
		return
	}

	dj.output.Success(name, resp, w)
}

func (dj *DojoHttp) dispatchInteractive(w http.ResponseWriter, r *http.Request) {
	input, err := util.ToStruct[*dto.InteractiveDispatchParams](r.Body)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	resp, err := dj.gateway.DispatchInteractive(*input)
	if err != nil {
		dj.output.Error(input, err.Error(), w)
		return
	}

	dj.output.Success(input, resp, w)
}

func (dj *DojoHttp) Handlers(r chi.Router) *DojoHttp {
	r.Route("/v1/dojo", func(r chi.Router) {
		r.Post("/", dj.dispatchMessage)
		r.Post("/flow", dj.createFlowHandler)
		r.Get("/flow/{name}", dj.getFlowHandler)
		r.Post("/message", dj.sendSimpleMessageHandler)
		r.Post("/template", dj.registerTempalteMeta)
		r.Post("/template/send-one", dj.sendOneTemplate)
		r.Get("/template/register", dj.modelRegisterTemplate)
		r.Get("/template/send", dj.modelTemplateByWrite)
		r.Post("/interactive", dj.createMsgInteractive)
		r.Get("/interactive", dj.interactiveMsgByName)
		r.Post("/interactive/dispatch", dj.dispatchInteractive)
		r.Get("/interactive/{id}", dj.interactiveMsgByID)

	})
	return dj
}
