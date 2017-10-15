package channel

import (
	"botwhatsapp/internal/app/whatsapp"
	"botwhatsapp/internal/app/whatsapp/usecases/dojo"
	"botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Dispatcher struct {
	wtsgate  whatsapp.Gateway
	dojogate dojo.Gateway
}

func NewDispatch(gate whatsapp.Gateway, dojo dojo.Gateway) *Dispatcher {
	return &Dispatcher{wtsgate: gate, dojogate: dojo}
}

var fluxos = map[string]string{
	"padrao": "Ola, escolha um fluxo\n\n1. fluxo de mensagem simples\n2. fluxo de mensagem interativas\n3. fluxo de mensagens dinamicas",
	"1":      "m1. Mensagem 1\nm2. Mensagem 2\nv. Voltar",
	"m1":     "M1 .............",
	"m2":     "M2 .............",
	"sair":   "Tchau ..........",
	"2":      "Fluxo2Mensagem",
	"3":      "Fluxo3Mensagem",
}

func (disp *Dispatcher) Dispatch(userData <-chan *UserData) {
	var wg sync.WaitGroup
	wg.Add(1)
	//flowExists := make(map[string]*dto.FlowData)
	go func() {
		defer wg.Done()
		for cc := range userData {
			//_, ok := flowExists[cc.MetaID]
			//if !ok {
			//	resp, err := disp.dojogate.DojoGetFlow(cc.MetaID, "init")
			//	if err != nil {
			//		fmt.Println("Erro ao obter fluxo do banco de dados:", err)
			//		return
			//	}
			//	flowExists[cc.MetaID] = resp
			//}
			//flow := flowExists[cc.MetaID]
			//_, ok := fluxos[cc.Payload]
			//if !ok {
			//	disp.SendSimpleMessage(cc.UserPhone, "Não entendi, qual dessas opções ?")
			//	disp.SendSimpleMessage(cc.UserPhone, fluxos["padrao"])
			//} else {
			//	disp.SendSimpleMessage(cc.UserPhone, fluxos[cc.Payload])
			//}

			fmt.Println("\n\n message: ", cc.Message.Messages[0].ID)
			disp.SendSimpleMessage(cc.UserPhone, cc.Message.Messages[0].ID, "teste")
		}
	}()

	go func() { wg.Wait() }()
}

func (disp *Dispatcher) workers() int {
	numWorkersStr := os.Getenv("NUMBER_WORKERS")
	numWorkers, err := strconv.Atoi(numWorkersStr)
	if err != nil || numWorkers <= 0 {
		numWorkers = 1
	}
	return numWorkers
}

func (disp *Dispatcher) SendSimpleMessage(to, mid, msg string) {
	txt := "text"
	simpleMessage := dojo.InputSendSimpleMessage{Message: msg, To: to, Type: &txt, MID: &mid}
	_, _ = disp.dojogate.SendSimpleMessage(&simpleMessage)
}

func (disp *Dispatcher) SendTemplateMessage(to, template string, params map[string]any) {
	payload := dto.DojoDispatchParams{
		Name:   &template,
		To:     &to,
		Next:   nil,
		Params: &params,
	}
	_, _ = disp.dojogate.Dispatch(&payload)
}

/**
f6605c7b899a2de18a227839e5c1543d
f6605c7b899a2de18a227839e5c1543d
f6605c7b899a2de18a227839e5c1543d
f6605c7b899a2de18a227839e5c1543d
f6605c7b899a2de18a227839e5c1543d
f6605c7b899a2de18a227839e5c1543d
f6605c7b899a2de18a227839e5c1543d
*/
