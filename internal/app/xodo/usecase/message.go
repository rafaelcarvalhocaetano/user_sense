package usecase

import (
	"botwhatsapp/internal/app/xodo/dto"
	"botwhatsapp/internal/interfaces/services"
	"errors"
	"fmt"
)

type Message struct {
	http services.WAGateway
}

func NewMessage(http services.WAGateway) *Message {
	return &Message{http: http}
}

func (msg *Message) SendMessage(data *dto.InputMessage) (*dto.OutputMessage, error) {
	if data.To == "" || len(data.To) != 13 {
		return nil, errors.New("invalid parameter")
	}

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"to":                data.To,
	}

	if data.Type == "document" {
		payload["type"] = "document"
		payload["document"] = map[string]interface{}{
			"caption": data.Caption,
			"link":    data.Link,
		}
	}

	if data.Type == "image" {
		payload["type"] = "image"
		payload["preview_url"] = true
		payload["context"] = map[string]interface{}{"message_id": data.MessageID}
		payload["image"] = map[string]interface{}{
			"link":    data.Link,
			"caption": data.Caption,
		}
	}

	if data.Type == "" || data.Type == "text" {
		payload["type"] = "text"
		payload["text"] = map[string]interface{}{
			"preview_url": true,
			"body":        data.Message,
		}
	}

	if data.Type == "contact" {
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

	_, err := msg.http.Send("messages", payload)
	if err != nil {
		return nil, errors.New("failed to send message")
	}

	msgTxt := fmt.Sprintf("message dispatch to %v", &data.To)
	return &dto.OutputMessage{Message: msgTxt}, nil
}
