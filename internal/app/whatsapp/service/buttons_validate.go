package service

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"fmt"
)

type ButtonsValidate struct{}

func NewButtonsValidate() *ButtonsValidate { return &ButtonsValidate{} }

// TODO: btn de copy não suporta o tipo de template UTILITY
// TODO: BTN quando o btn for do tipo url sub_type deve ser url
// TODO: BTN listagem de btn deve ter o index
// ACTION, COUPON_CODE, CURRENCY, DATE_TIME, DOCUMENT, EXPIRATION_TIME_MS, IMAGE,
// LIMITED_TIME_OFFER, LOCATION, ORDER_STATUS, PAYLOAD, PRODUCT, TEXT,
// VIDEO, WEBVIEW_INTERACTION, WEBVIEW_PRESENTATION
func (b *ButtonsValidate) CheckParameters(data map[string]any) (*[]entities.TemplateComponent, error) {
	templatesComponent := make([]entities.TemplateComponent, 0)

	buttons, ok := data["buttons"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("'buttons' key is missing or not of type []interface{}")
	}

	for i, btn := range buttons {
		btnMap, ok := btn.(map[string]interface{})
		if !ok {
			continue
		}

		examples, ok := btnMap["example"].([]interface{})
		if !ok || examples == nil {
			continue
		}

		index := fmt.Sprintf("%v", i)
		subType, ok := btnMap["type"].(string)
		if !ok {
			continue
		}

		templateComponent := entities.TemplateComponent{
			Type:       "button",
			SubType:    &subType,
			Index:      &index,
			Parameters: make([]*entities.TemplateParameter, 0),
		}

		for range examples {
			q := "[here]"
			paramButton := entities.TemplateParameter{Type: "text", Text: &q}
			templateComponent.Parameters = append(templateComponent.Parameters, &paramButton)
		}

		templatesComponent = append(templatesComponent, templateComponent)
	}

	return &templatesComponent, nil
}
