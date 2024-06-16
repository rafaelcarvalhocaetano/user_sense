package channel

import (
	xodo2 "botwhatsapp/internal/app/xodo"
	"botwhatsapp/internal/app/xodo/dto"
	"botwhatsapp/internal/interfaces/webhooks/model"
	"os"
	"strconv"
	"sync"
)

type ChannelFlowXodo struct {
	xodo xodo2.Gateway
}

func NewChannelFlowXodo(gate xodo2.Gateway) *ChannelFlowXodo {
	return &ChannelFlowXodo{xodo: gate}
}

func (cx *ChannelFlowXodo) ChannelFlowXodo(
	userData <-chan *UserData,
	statuses <-chan model.Status,
	broker <-chan model.Channel,
) {
	var wg sync.WaitGroup
	wg.Add(1)

	processMap := make(map[string]map[string]string)

	//go func() {
	//	defer wg.Done()
	//	for s := range statuses {
	//		//_, _ = pp.Println(s)
	//		if s.Status == "sent" {
	//			processMap[*s.RecipientID] = map[string]string{"step": "continue"}
	//		}
	//	}
	//}()

	go func() {
		defer wg.Done()
		for b := range broker {
			processMap[b.PhoneNumber] = map[string]string{"step": "continue"}
		}
	}()

	go func() {
		defer wg.Done()
		for user := range userData {

			processor, ok := processMap[user.UserPhone]
			if processor["step"] == "stop" || !ok {
				simpleMessage := dto.InputMessage{
					To:            user.UserPhone,
					Type:          "contact",
					ContactNumber: "556281680703",
					ContactName:   "STP CLUB",
				}
				cx.sendSimpleMessage(user.UserPhone, "Converse com nosso suporte, pelo contato: ")
				_, _ = cx.xodo.SendMessage(&simpleMessage)
			}

			if processor["step"] == "continue" {
				switch user.Payload {
				case "1", "2", "3", "4", "5":
					msg := "Obrigado por nos avaliar! Visite-nos novamente em: \n\nhttps://www.instagram.com/reel/C4tBLGeuON2/?igsh=YW5ta3MxMDFjczF1\n\nSe preferir, entre em contato com nosso suporte:"
					cx.sendSimpleMessage(user.UserPhone, msg)
					//go func() {
					//	input := dto.InputRate{PhoneNumber: user.UserPhone}
					//	_, _ = cx.xodo.Mkt(input)
					//	processor["step"] = "stop"
					//}()
					cx.Suport(user.UserPhone)
					continue
				default:
					cx.sendSimpleMessage(user.UserPhone, "Desculpe, não conseguimos entender. Por favor, classifique de 1 a 5.")
				}
			}

		}
	}()
	go func() { wg.Wait() }()
}

func (cx *ChannelFlowXodo) Suport(p string) {
	simpleMessage := dto.InputMessage{
		To:            p,
		Type:          "contact",
		ContactNumber: "556281680703",
		ContactName:   "STP CLUB",
	}
	_, _ = cx.xodo.SendMessage(&simpleMessage)

}

func (cx *ChannelFlowXodo) workers() int {
	numWorkersStr := os.Getenv("NUMBER_WORKERS")
	numWorkers, err := strconv.Atoi(numWorkersStr)
	if err != nil || numWorkers <= 0 {
		numWorkers = 1
	}
	return numWorkers
}

func (cx *ChannelFlowXodo) sendSimpleMessage(to, msg string) {
	simpleMessage := dto.InputMessage{Message: msg, To: to, Type: "text"}
	_, _ = cx.xodo.SendMessage(&simpleMessage)
}
