package service

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"errors"
	"fmt"
)

type BodyValidate struct{}

func NewBodyValidate() *BodyValidate { return &BodyValidate{} }

func (t *BodyValidate) CheckParameters(data map[string]interface{}) ([]*entities.TemplateParameter, error) {
	examples := MapperExemple(data)
	if examples == nil {
		return nil, nil
	}

	params := make([]*entities.TemplateParameter, 0)
	for _, example := range examples {
		itemList, ok := example.([]interface{})
		if !ok {
			return nil, errors.New("invalid example format")
		}

		for _, item := range itemList {
			param := entities.TemplateParameter{}
			paramType := CheckType(item)
			switch paramType {
			case string(entities.MessageTypeText):
				param.Type = string(entities.MessageTypeText)
				text := "[here]"
				param.Text = &text
			case string(entities.MessageTypeCurrency):
				param.Type = string(entities.MessageTypeCurrency)
				currencyValue, err := ParseCurrency(item.(string))
				if err != nil {
					return nil, fmt.Errorf("failed to parse currency: %w", err)
				}
				param.Currency = &entities.TemplateCurrency{
					FallbackValue: "[here]",
					Code:          "BRL",
					Amount1000:    currencyValue,
				}
			case string(entities.MessageTypeDate):
				param.Type = string(entities.MessageTypeDate)
				param.DateTime = &entities.TemplateDate{
					FallbackValue: "[here]",
					DayOfWeek:     0,
					Year:          0,
					Month:         0,
					DayOfMonth:    0,
					Hour:          0,
					Minute:        0,
					Calendar:      "GREGORIAN",
				}
			case string(entities.MessageTypeImage):
				param.Type = string(entities.MessageTypeImage)
				link := "[here]"
				param.Image = &entities.TemplateImage{Link: link}
			case string(entities.MessageTypeDocument):
				param.Type = string(entities.MessageTypeDocument)
				link := "[here]"
				param.Image = &entities.TemplateImage{Link: link}
			default:
				return nil, fmt.Errorf("unsupported message type: %v", paramType)
			}

			params = append(params, &param)
		}
	}

	return params, nil
}

//func (t *BodyValidate) CheckParamaters(data map[string]any) ([]*dto.TemplateParameter, error) {
//	p1 := "[here]"
//	examples := MapperExemple(data)
//	if examples != nil {
//		var params []*dto.TemplateParameter
//		for _, item := range examples {
//			var param dto.TemplateParameter
//			itemList := item.([]any)
//			for _, d := range itemList {
//				checkType := CheckType(d)
//				if checkType == string(entities.MessageTypeText) {
//					param.Type = string(entities.MessageTypeText)
//					param.Text = &p1
//				}
//				if checkType == string(entities.MessageTypeCurrency) {
//					param.Type = string(entities.MessageTypeCurrency)
//					valorLimpo := strings.Replace(item.(string), "R$", "", 1)
//					valorLimpo = strings.Replace(valorLimpo, ".", "", 1)
//					valorInt, _ := strconv.Atoi(valorLimpo)
//					param.Currency = &dto.TemplateCurrency{
//						FallbackValue: p1,
//						Code:          "BRL",
//						Amount1000:    valorInt,
//					}
//				}
//				if checkType == string(entities.MessageTypeDate) {
//					param.Type = string(entities.MessageTypeDate)
//					param.DateTime = &dto.TemplateDate{
//						FallbackValue: p1,
//						DayOfWeek:     0,
//						Year:          0,
//						Month:         0,
//						DayOfMonth:    0,
//						Hour:          0,
//						Minute:        0,
//						Calendar:      "GREGORIAN",
//					}
//				}
//				if checkType == string(entities.MessageTypeImage) {
//					param.Type = string(entities.MessageTypeImage)
//					param.Image = &dto.TemplateImage{Link: p1}
//				}
//				if checkType == string(entities.MessageTypeDocument) {
//					param.Type = string(entities.MessageTypeDocument)
//					param.Image = &dto.TemplateImage{Link: p1}
//				}
//				params = append(params, &param)
//			}
//		}
//
//		return params, nil
//	}
//
//	return nil, nil
//}
