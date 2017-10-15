package dojo

import (
	"botwhatsapp/internal/interfaces/services"
)

type SendOneTemplate struct {
	gateway    services.WTAGateway
	saveStatus SaveStatusUserMessage
}

func NewSendOneTemplate(gateway services.WTAGateway, save SaveStatusUserMessage) *SendOneTemplate {
	return &SendOneTemplate{gateway: gateway, saveStatus: save}
}

func (send *SendOneTemplate) SendOneTemplate(data map[string]any) (*map[string]any, error) {
	response := make(map[string]any)
	resp, err := send.gateway.Send("messages", data)
	if err != nil {
		response["error"] = err.Error()
		return nil, err
	}

	phone := data["to"].(string)
	err = send.saveStatus.SaveStatusUserMessage(phone, resp)
	if err != nil {
		response["error"] = err.Error()
		return &response, nil
	}

	response["success"] = resp
	return &response, nil
}
