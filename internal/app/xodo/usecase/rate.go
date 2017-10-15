package usecase

import (
	"botwhatsapp/internal/interfaces/services"
	"fmt"
)

type objectData map[string]interface{}

type Rate struct {
	gateway services.WTAGateway
}

func NewRate(gateway services.WTAGateway) *Rate {
	return &Rate{gateway: gateway}
}

func (r *Rate) Rate() (any, error) {
	var rates []objectData
	start := "⭐"

	for i := 0; i < 5; i++ {
		rates = append(rates, objectData{
			"id":          fmt.Sprintf("%d", i),
			"title":       start,
			"description": "",
		})
		start += "⭐"
	}

	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"type":              "interactive",
		"to":                "5511966896414",
		"interactive": map[string]interface{}{
			"type": "list",
			"header": map[string]interface{}{
				"type": "image",
				"text": "",
				"image": map[string]interface{}{
					"link": "https://media.guiame.com.br/archives/2017/12/28/2306764763-filme-eu-so-posso-imaginar-01.png",
				},
			},
			"body": objectData{
				"text": "xxxxxxxxxxxxxx",
			},
			"action": objectData{
				"button": "Avaliar",
				"sections": []interface{}{
					objectData{
						"title": "AVX",
						"rows":  rates,
					},
				},
			},
		},
	}

	resp, err := r.gateway.Send("messages", payload)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
