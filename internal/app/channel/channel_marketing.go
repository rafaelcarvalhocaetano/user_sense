package channel

import (
	dto2 "botwhatsapp/internal/app/channel/dto"
	xodo2 "botwhatsapp/internal/app/xodo"
	"botwhatsapp/internal/app/xodo/dto"
	"botwhatsapp/internal/interfaces/webhooks/model"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type ChannelMkt struct {
	xodo xodo2.Gateway
}

func NewChannelMkt(gate xodo2.Gateway) *ChannelMkt {
	return &ChannelMkt{xodo: gate}
}

func (cx *ChannelMkt) Flow(messages <-chan *dto2.ChannelDTO, statuses <-chan model.Status, broker <-chan model.Channel) {
	var wg sync.WaitGroup
	wg.Add(1)

	processMap := make(map[string]map[string]string)

	go func() {
		defer wg.Done()
		for s := range statuses {
			fmt.Println("statuses", s)
		}
	}()

	go func() {
		defer wg.Done()
		for b := range broker {
			if b.Status {
				processMap[b.PhoneNumber] = map[string]string{"step": "continue"}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for m := range messages {
			processor, ok := processMap[m.UserPhone]
			if processor["step"] == "stop" || !ok {
				simpleMessage := dto.InputMessage{
					To:            m.UserPhone,
					Type:          "contact",
					ContactNumber: "556281680703",
					ContactName:   "STP CLUB",
				}
				cx.messageTxt(m.UserPhone, "Converse com nosso suporte, pelo contato: ")
				_, _ = cx.xodo.SendMessage(&simpleMessage)
			}

			if processor["step"] == "continue" {
				switch m.Payload {
				case "1", "2", "3", "4", "5":
					cx.messageEnded(m.UserPhone)
					time.Sleep(time.Second * 1)
					cx.suport(m.UserPhone)
					processor["step"] = "stop"
					continue
				default:
					cx.messageTxt(m.UserPhone, "Desculpe, não conseguimos entender. Por favor, classifique de 1 a 5.")
				}
			}

		}
	}()
	go func() { wg.Wait() }()
}

func (cx *ChannelMkt) messageEnded(p string) {
	image := "https://github.com/rafaelcarvalhocaetano/meetup/blob/master/seja.png?raw=true"
	insta := "https://www.instagram.com/reel/C4tBLGeuON2/?igsh=YW5ta3MxMDFjczF1"
	suport := "\n\nSe preferir, entre em contato com nosso suporte:"
	msg := "Obrigado por nos avaliar! Visite-nos novamente em:\n\n " + insta + suport
	simpleMessage := dto.InputMessage{
		To:      p,
		Type:    "image",
		Link:    image,
		Caption: msg,
	}
	_, _ = cx.xodo.SendMessage(&simpleMessage)
}

func (cx *ChannelMkt) suport(p string) {
	simpleMessage := dto.InputMessage{
		To:            p,
		Type:          "contact",
		ContactNumber: "556281680703",
		ContactName:   "STP CLUB",
	}
	_, _ = cx.xodo.SendMessage(&simpleMessage)
}

func (cx *ChannelMkt) workers() int {
	numWorkersStr := os.Getenv("NUMBER_WORKERS")
	numWorkers, err := strconv.Atoi(numWorkersStr)
	if err != nil || numWorkers <= 0 {
		numWorkers = 1
	}

	return numWorkers
}

func (cx *ChannelMkt) messageTxt(to, msg string) {
	simpleMessage := dto.InputMessage{Message: msg, To: to, Type: "text"}
	_, _ = cx.xodo.SendMessage(&simpleMessage)
}
