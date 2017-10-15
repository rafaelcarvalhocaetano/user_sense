package dojo

import (
	"botwhatsapp/internal/interfaces/services"
	"errors"
	"fmt"
)

type SendSimpleMessage struct {
	dispatch services.WTAGateway
}

type outputSendSimpleMessage struct {
	Message string `json:"body"`
}

type InputSendSimpleMessage struct {
	To            string  `json:"to"`
	Message       string  `json:"message"`
	Type          *string `json:"type,omitempty"`
	ContactNumber *string `json:"contact_number,omitempty"`
	ContactName   *string `json:"contact_name,omitempty"`
	Caption       *string `json:"caption,omitempty"`
	MID           *string `json:"mid,omitempty"`
}

func NewSendSimpleMessage(d services.WTAGateway) *SendSimpleMessage {
	return &SendSimpleMessage{dispatch: d}
}

func (sm *SendSimpleMessage) SendSimpleMessage(data *InputSendSimpleMessage) (*outputSendSimpleMessage, error) {
	if data.To == "" || len(data.To) != 13 {
		return nil, errors.New("invalid parameter")
	}

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                data.To,
	}

	if *data.Type == "document" {
		payload["type"] = "document"
		payload["document"] = map[string]interface{}{
			"caption": data.Caption,
			"link":    data.Message,
		}
	}

	if *data.Type == "image" {
		payload["type"] = "image"
		payload["image"] = map[string]interface{}{
			"preview_url": true,
			"body":        data.Message,
		}
	}

	if data.Type == nil || *data.Type == "text" {
		payload["type"] = "text"
		//payload["context"] = map[string]any{
		//	"message_id": data.MID,
		//}
		payload["text"] = map[string]interface{}{
			"preview_url": true,
			"body":        data.Message,
		}
	}

	if *data.Type == "contact" {
		payload["type"] = "contacts"
		payload["contacts"] = []map[string]interface{}{
			{
				"name": map[string]interface{}{
					"first_name":     data.ContactName,
					"last_name":      "Smith",
					"formatted_name": data.ContactName,
				},
				"phones": []map[string]interface{}{
					{
						"phone": data.ContactNumber,
						"wa_id": data.ContactNumber,
						"type":  "WORK",
					},
				},
			},
		}
	}

	_, err := sm.dispatch.Send("messages", payload)
	if err != nil {
		return nil, errors.New("failed to send message")
	}

	msg := fmt.Sprintf("message dispatch to %v", &data.To)
	return &outputSendSimpleMessage{Message: msg}, nil
}
