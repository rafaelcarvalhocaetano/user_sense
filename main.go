package main

import (
	"botwhatsapp/internal/app/xodo"
	"botwhatsapp/internal/interfaces/http"
	"botwhatsapp/internal/interfaces/services/whatsapp"
	"botwhatsapp/internal/interfaces/web"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"botwhatsapp/internal/app/channel"
	"botwhatsapp/internal/infra/drivers"
	"botwhatsapp/internal/interfaces/services"
	"botwhatsapp/internal/interfaces/webhooks"
	"botwhatsapp/internal/interfaces/webhooks/model"
)

func main() {
	_ = godotenv.Load(".env")

	// TODO: loggers - zap-fx
	cloudwatchDriver := drivers.NewCloudwatchDriver()
	logger := cloudwatchDriver.InitializeLoggers()
	logDriver := drivers.NewLoggerCloudwatch(logger)

	// TODO: adapters
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))

	channelManager := make(chan *model.WebhookData)
	dispather := make(chan *channel.UserData)
	templateChannel := make(chan model.Status)
	statusChannel := make(chan model.Channel)

	// TODO: service httpClient
	sendMessages := whatsapp.NewSendMessage(logDriver)
	templates := whatsapp.NewTemplateService(logDriver)
	register := whatsapp.NewRegisteTemplateHttp(logDriver)
	deleteTemplate := whatsapp.NewDeleteTemplateHttp(logDriver)

	httpGatway := services.WTAGateway{
		SendMessagesHttp:      sendMessages,
		RegisterHttp:          register,
		GetTemplateGateway:    templates,
		DeleteTemplateGateway: deleteTemplate,
	}

	// TODO: Xodo
	xd := xodo.New().Main(httpGatway, statusChannel)
	http.NewXodoHttp(logDriver, *xd).Handlers(r)

	// TODO: webhooks
	webhooks.NewWhatsApp(logDriver, httpGatway).Handler(r, channelManager)
	channel.NewMessageChannel().Main(channelManager, dispather, templateChannel)

	channel.NewChannelFlowXodo(*xd).ChannelFlowXodo(dispather, templateChannel, statusChannel)
	//channel.NewChannelStatusMessage().ChannelStatusMessage(templateChannel)

	// TODO: web-server
	web.NewServer(r)
}
