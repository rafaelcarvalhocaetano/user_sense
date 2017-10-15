package message

import (
	"botwhatsapp/internal/app/whatsapp/usecases/message/dto"
	"botwhatsapp/internal/interfaces/services"
	"errors"
	"fmt"
)

type SimpleMessage struct {
	dispatch services.WTAGateway
}

func NewSimpleMessage(d services.WTAGateway) *SimpleMessage {
	return &SimpleMessage{dispatch: d}
}

func (sm *SimpleMessage) SimpleMessage(data *dto.MessageData) (*dto.OutputSimpleMessage, error) {
	if data.To == "" || len(data.To) != 13 {
		return nil, errors.New("invalid parameter")
	}

	data.MessagingProduct = "whatsapp"
	_, err := sm.dispatch.Send("messages", data)
	if err != nil {
		return nil, errors.New("failed to send message")
	}

	msg := fmt.Sprintf("message dispatch to %v", &data.To)
	return &dto.OutputSimpleMessage{Message: msg}, nil
}
