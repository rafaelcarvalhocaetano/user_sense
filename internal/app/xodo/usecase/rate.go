package usecase

import (
	"botwhatsapp/internal/app/xodo/dto"
	"botwhatsapp/internal/interfaces/services"
)

type objectData map[string]interface{}

type Rate struct {
	gateway services.WTAGateway
}

func NewRate(gateway services.WTAGateway) *Rate {
	return &Rate{gateway: gateway}
}

func (r *Rate) Rate(input dto.Input) (any, error) {
	//var rates []objectData
	//start := "⭐"

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
								"link": "https://s2-techtudo.glbimg.com/SSAPhiaAy_zLTOu3Tr3ZKu2H5vg=/0x0:1024x609/888x0/smart/filters:strip_icc()/i.s3.glbimg.com/v1/AUTH_08fbf48bc0524877943fe86e43087e7a/internal_photos/bs/2022/c/u/15eppqSmeTdHkoAKM0Uw/dall-e-2.jpg",
							},
						},
					},
				},
			},
		},
	}

	// SEND TEMPLATE IMAGE
	resp, err := r.gateway.Send("messages", templatePayload)
	if err != nil {
		return nil, err
	}

	//fmt.Println("\n\n resp: ", resp)
	//
	//for i := 0; i < 5; i++ {
	//	rates = append(rates, objectData{
	//		"id":          fmt.Sprintf("%d", i),
	//		"title":       start,
	//		"description": "",
	//	})
	//	start += "⭐"
	//}
	//
	//payload := map[string]interface{}{
	//	"messaging_product": "whatsapp",
	//	"recipient_type":    "individual",
	//	"type":              "interactive",
	//	"to":                input.PhoneNumber,
	//	"interactive": map[string]interface{}{
	//		"type": "list",
	//		"body": objectData{
	//			"text": "Avaliar",
	//		},
	//		"action": objectData{
	//			"button": "Avaliar",
	//			"sections": []interface{}{
	//				objectData{
	//					"title": "avaliar",
	//					"rows":  rates,
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//resp, err = r.gateway.Send("messages", payload)
	//if err != nil {
	//	return nil, err
	//}

	return resp, nil
}
