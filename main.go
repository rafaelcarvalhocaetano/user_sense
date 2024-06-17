package main

import (
	"botwhatsapp/internal/app/channel"
	"botwhatsapp/internal/app/channel/dto"
	"botwhatsapp/internal/app/xodo"
	"botwhatsapp/internal/infra/drivers"
	"botwhatsapp/internal/interfaces/http"
	"botwhatsapp/internal/interfaces/services"
	"botwhatsapp/internal/interfaces/web"
	"botwhatsapp/internal/interfaces/webhooks"
	"botwhatsapp/internal/interfaces/webhooks/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
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

	webhookChannel := make(chan *model.WebhookData)
	messageChannel := make(chan *dto.ChannelDTO)
	statusesChannel := make(chan model.Status)
	sendChannel := make(chan model.Channel)

	// TODO: service httpClient
	sendMessages := services.NewSendMessage(logDriver)
	httpGatway := services.WAGateway{SendMessagesHttp: sendMessages}

	// TODO: Xodo
	xd := xodo.New().Main(httpGatway, sendChannel)
	http.NewXodoHttp(logDriver, *xd).Handlers(r)

	// TODO: webhooks
	webhooks.NewWhatsapp(logDriver, httpGatway).Handler(r, webhookChannel)
	channel.NewMessageChannel().Main(webhookChannel, messageChannel, statusesChannel)

	channel.NewChannelMkt(*xd).Flow(messageChannel, statusesChannel, sendChannel)

	// TODO: web-server
	web.NewServer(r)
	//appengine.Main()
}
