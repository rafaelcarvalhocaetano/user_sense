package webhooks

import (
	"botwhatsapp/internal/infra/ports"
	"botwhatsapp/internal/interfaces/services"
	"botwhatsapp/internal/interfaces/webhooks/model"
	"botwhatsapp/internal/pkg/response"
	"botwhatsapp/util"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
)

type Whatsapp struct {
	log        ports.Logger
	output     *response.Response
	httpClinet services.WAGateway
}

func NewWhatsapp(log ports.Logger, httpClinet services.WAGateway) *Whatsapp {
	return &Whatsapp{
		output:     response.NewResponse("webhook", log),
		log:        log,
		httpClinet: httpClinet,
	}
}

func (wt *Whatsapp) webhookAuthentication(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("hub.challenge")
	_, err := w.Write([]byte(challenge))
	if err != nil {
		return
	}
}

func (wt *Whatsapp) messages(w http.ResponseWriter, r *http.Request, dataChan chan<- *model.WebhookData) {
	input, _ := util.ToStruct[model.WebhookData](r.Body)

	go func() {
		dataChan <- input
	}()

	//_, _ = pp.Println(input)

	body, _ := io.ReadAll(r.Body)
	log.Println(string(body))
	wt.log.Debug("webhook", body, http.StatusOK)

}

func (wt *Whatsapp) Handler(r chi.Router, dataChan chan<- *model.WebhookData) *Whatsapp {
	r.Get("/whatsapp", wt.webhookAuthentication)
	r.Post("/whatsapp", func(w http.ResponseWriter, r *http.Request) {
		wt.messages(w, r, dataChan)
	})
	return wt
}
