package usecase

import (
	"botwhatsapp/internal/app/xodo/dto"
	"botwhatsapp/internal/interfaces/services"
	"botwhatsapp/internal/interfaces/webhooks/model"
)

type objectData map[string]interface{}

type Rate struct {
	gateway services.WAGateway
	x       chan model.Channel
}

func NewRate(gateway services.WAGateway, x chan model.Channel) *Rate {
	return &Rate{gateway: gateway, x: x}
}

func (r *Rate) Rate(input dto.InputRate) (*string, error) {
	var imageUrl = "https://github.com/rafaelcarvalhocaetano/meetup/blob/master/aval.png?raw=true"

	templatePayload := objectData{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                input.PhoneNumber,
		"type":              "template",
		"template": objectData{
			"name":     "xodo_flow",
			"language": map[string]string{"code": "pt_BR"},
			"components": []objectData{
				{
					"type": "header",
					"parameters": []objectData{
						{
							"type": "image",
							"image": objectData{
								"link": imageUrl,
							},
						},
					},
				},
			},
		},
	}

	respID, err := r.gateway.Send("messages", templatePayload)
	if err != nil {
		return nil, err
	}

	r.x <- model.Channel{
		PhoneNumber: input.PhoneNumber,
		Status:      true,
	}

	return respID, nil
}
