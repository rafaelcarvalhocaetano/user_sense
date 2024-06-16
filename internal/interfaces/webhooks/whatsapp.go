package webhooks

import (
	"botwhatsapp/internal/infra/ports"
	"botwhatsapp/internal/interfaces/services"
	"botwhatsapp/internal/interfaces/webhooks/model"
	"botwhatsapp/internal/pkg/response"
	"botwhatsapp/util"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type WhatsApp struct {
	log        ports.Logger
	output     *response.Response
	httpClinet services.WTAGateway
}

func NewWhatsApp(log ports.Logger, httpClinet services.WTAGateway) *WhatsApp {
	return &WhatsApp{
		output:     response.NewResponse("webhook", log),
		log:        log,
		httpClinet: httpClinet,
	}
}

func (wt *WhatsApp) authentication(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aqui \n\n")
	challenge := r.URL.Query().Get("hub.challenge")
	_, err := w.Write([]byte(challenge))
	if err != nil {
		return
	}
}

func (wt *WhatsApp) messages(w http.ResponseWriter, r *http.Request, dataChan chan<- *model.WebhookData) {
	input, err := util.ToStruct[model.WebhookData](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	//_, _ = pp.Println(input)
	go func() {
		dataChan <- input
	}()

	w.WriteHeader(http.StatusOK)
	//_, err = w.Write(wt.output.Success(input, "success"))
	//if err != nil {
	//	return
	//}
}

func (wt *WhatsApp) Handler(r chi.Router, dataChan chan<- *model.WebhookData) *WhatsApp {
	r.Get("/whatsapp", wt.authentication)
	r.Post("/whatsapp", func(w http.ResponseWriter, r *http.Request) {
		wt.messages(w, r, dataChan)
	})
	return wt
}

//func validateStruct(data map[string]any) *string {
//	entryList := data["entry"].([]any)
//	for _, entry := range entryList {
//		entryMap := entry.(map[string]any)
//		changes := entryMap["changes"].([]any)
//		for _, change := range changes {
//			_, ok := change.(map[string]any)["messages"]
//			if ok {
//				m := "mensage"
//				return &m
//			}
//			_, ok = change.(map[string]any)["statuses"]
//			if ok {
//				m := "statuses"
//				return &m
//			}
//		}
//	}
//	return nil
//}
