package usecase

import (
	"botwhatsapp/internal/app/xodo/dto"
	"botwhatsapp/internal/interfaces/services"
)

type Marketing struct {
	gateway services.WTAGateway
}

func NewMarketing(gateway services.WTAGateway) *Marketing {
	return &Marketing{gateway: gateway}
}

func (r *Marketing) Mkt(input dto.InputRate) (*string, error) {
	var imageUrl = "https://github.com/rafaelcarvalhocaetano/meetup/blob/master/seja.png?raw=true"

	templatePayload := objectData{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                input.PhoneNumber,
		"type":              "template",
		"template": objectData{
			"name":     "xodo_propaganda",
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

	return respID, nil
}
