package channel

import (
	xodo2 "botwhatsapp/internal/app/xodo"
	"botwhatsapp/internal/interfaces/webhooks/model"
	"fmt"
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

var aval = "https://github.com/rafaelcarvalhocaetano/wa_xodo/blob/master/assets/aval.jpeg?raw=true"
var membro = "https://github.com/rafaelcarvalhocaetano/meetup/blob/master/membro.png?raw=true"
var xodo = "https://media.licdn.com/dms/image/D4D0BAQEe8aamu2cmTg/company-logo_200_200/0/1711134503656/xodoapp_logo?e=2147483647&v=beta&t=wWm6gg64bolgEo6jFctyn2gdt9jzjiPFVHotyS0fzTk"

// var sac = "https://github.com/rafaelcarvalhocaetano/meetup/blob/master/sac.png?raw=true"

func (disp *ChannelFlowXodo) ChannelFlowXodo(userData <-chan *UserData, statuses <-chan model.Status) {
	var wg sync.WaitGroup
	wg.Add(1)

	processMap := make(map[string]map[string]string)

	go func() {
		defer wg.Done()
		for s := range statuses {
			if processMap[*s.RecipientID]["id"] == s.ID && s.Status != "accepted" {
				processMap[*s.RecipientID]["status"] = "open"
			}
		}
	}()

	go func() {
		defer wg.Done()
		for user := range userData {

			fmt.Println("user: ", user)

			//if _, exists := processMap[user.UserPhone]; !exists {
			//	processMap[user.UserPhone] = map[string]string{
			//		"step":   "",
			//		"id":     "",
			//		"status": "",
			//	}
			//}

			//if processMap[user.UserPhone]["status"] == "closed" {
			//	go disp.sendSimpleMessage(user.UserPhone, "Aguarde, estamos processando seu voucher ...")
			//	continue
			//}
			//
			//if processMap[user.UserPhone]["step"] == "1" {
			//	processMap[user.UserPhone]["evaluation"] = user.Payload
			//	switch user.Payload {
			//	case "1":
			//		disp.sendSimpleMessage(user.UserPhone, "Poxa, que pena. Vamos melhorar.")
			//		continue
			//	case "2", "3":
			//		disp.sendSimpleMessage(user.UserPhone, "Obrigado. Continuaremos a melhorar.")
			//		continue
			//	}
			//}
			//
			//switch strings.ToUpper(user.Payload) {
			//case "AVALIAR ABASTECIMENTO":
			//	processMap[user.UserPhone]["step"] = "1"
			//	p := map[string]any{
			//		"button_1_id":    "1",
			//		"button_1_title": "Ruim",
			//		"button_2_id":    "2",
			//		"button_2_title": "Bom",
			//		"button_3_id":    "3",
			//		"button_3_title": "Excelente",
			//	}
			//	disp.sendInteractiveMessage(user.UserPhone, "xodo_avaliacao", aval, "", "Qual sua avaliação para esse atendimento?", p)
			//case "TORNE-SE MEMBRO STP CLUB":
			//	processMap[user.UserPhone]["step"] = "2"
			//	p := map[string]interface{}{
			//		"param_header_1": membro,
			//		"param_next_1":   "rafaelcarvalhocaetano",
			//	}
			//	disp.sendSimpleMessage(user.UserPhone, "Legal, vamos gerar um voucher, aguarde ...")
			//	go func() {
			//		stt, err := disp.sendTemplateMessage(user.UserPhone, "xodo_qrcode", p)
			//		if err != nil {
			//			disp.sendSimpleMessage(user.UserPhone, "Algo de errado, tente novamente")
			//			return
			//		}
			//
			//		processMap[user.UserPhone]["id"] = stt.ID
			//		processMap[user.UserPhone]["status"] = "closed"
			//	}()
			//
			//case "SAC":
			//	processMap[user.UserPhone]["step"] = "3"
			//	disp.sendContact(user.UserPhone, "556281680703", "STP CLUB")
			//default:
			//	processMap[user.UserPhone] = map[string]string{
			//		"step":   "4",
			//		"id":     "",
			//		"status": "",
			//	}
			//	p := map[string]interface{}{"param_header_1": xodo}
			//	_, _ = disp.sendTemplateMessage(user.UserPhone, "xodo_welcome", p)
			//}
		}
	}()
	go func() { wg.Wait() }()
}

func (disp *ChannelFlowXodo) sendInteractiveMessage(phone, iname, header, fo, bo string, params map[string]any) {
	//interactiveParams := dto.InteractiveDispatchParams{
	//	Name:   &iname,
	//	To:     &phone,
	//	Header: &header,
	//	Body:   &bo,
	//	Footer: &fo,
	//	Params: params,
	//}
	//_, _ = disp.dojogate.DispatchInteractive(&interactiveParams)
}

func (disp *ChannelFlowXodo) workers() int {
	numWorkersStr := os.Getenv("NUMBER_WORKERS")
	numWorkers, err := strconv.Atoi(numWorkersStr)
	if err != nil || numWorkers <= 0 {
		numWorkers = 1
	}
	return numWorkers
}

func (disp *ChannelFlowXodo) sendSimpleMessage(to, msg string) {
	//t := "text"
	//simpleMessage := dojo.InputSendSimpleMessage{Message: msg, To: to, Type: &t}
	//_, _ = disp.dojogate.SendSimpleMessage(&simpleMessage)
}

func (disp *ChannelFlowXodo) sendContact(to, cnumber, cname string) {
	//t := "contact"
	//simpleMessage := dojo.InputSendSimpleMessage{
	//	To:            to,
	//	Type:          &t,
	//	ContactNumber: &cnumber,
	//	ContactName:   &cname,
	//}
	//_, _ = disp.dojogate.SendSimpleMessage(&simpleMessage)
}

type Stt struct {
	ID     string
	Status string
}

func (disp *ChannelFlowXodo) sendTemplateMessage(to, template string, params map[string]any) (*Stt, error) {
	//payload := dto.DojoDispatchParams{
	//	Name:   &template,
	//	To:     &to,
	//	Params: &params,
	//}

	//response, err := disp.dojogate.Dispatch(&payload)
	//if err != nil {
	//	return nil, err
	//}
	//resp, _ := response["resp"].(map[string]interface{})
	//messages, _ := resp["messages"].([]interface{})
	//message, _ := messages[0].(map[string]interface{})
	//id, _ := message["id"].(string)
	//status, _ := message["message_status"].(string)
	//fmt.Println("\n\n\ninfodata", id, " status ", status)
	//sttus := Stt{
	//	ID:     id,
	//	Status: status,
	//}
	//return &sttus, nil

	return nil, nil
}
