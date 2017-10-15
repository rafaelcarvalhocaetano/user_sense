package service

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"strconv"
	"strings"
)

type FooterValidate struct{}

func NewFooterValidate() *FooterValidate { return &FooterValidate{} }

func (f *FooterValidate) CheckParameters(data map[string]any) ([]*entities.TemplateParameter, error) {
	p1 := "[here]"
	examples := MapperExemple(data)
	if examples != nil {
		var params []*entities.TemplateParameter
		for _, item := range examples {
			var param entities.TemplateParameter
			itemList := item.([]any)
			for _, d := range itemList {
				checkType := CheckType(d)
				if checkType == string(entities.MessageTypeText) {
					param.Type = string(entities.MessageTypeText)
					param.Text = &p1
				}
				if checkType == string(entities.MessageTypeCurrency) {
					param.Type = string(entities.MessageTypeCurrency)
					valorLimpo := strings.Replace(item.(string), "R$", "", 1)
					valorLimpo = strings.Replace(valorLimpo, ".", "", 1)
					valorInt, _ := strconv.Atoi(valorLimpo)
					param.Currency = &entities.TemplateCurrency{
						FallbackValue: p1,
						Code:          "BRL",
						Amount1000:    valorInt,
					}
				}
				if checkType == string(entities.MessageTypeDate) {
					param.Type = string(entities.MessageTypeDate)
					param.DateTime = &entities.TemplateDate{
						FallbackValue: p1,
						DayOfWeek:     0,
						Year:          0,
						Month:         0,
						DayOfMonth:    0,
						Hour:          0,
						Minute:        0,
						Calendar:      "GREGORIAN",
					}
				}
				if checkType == string(entities.MessageTypeImage) {
					param.Type = string(entities.MessageTypeImage)
					param.Image = &entities.TemplateImage{Link: p1}
				}
				if checkType == string(entities.MessageTypeDocument) {
					param.Type = string(entities.MessageTypeDocument)
					param.Image = &entities.TemplateImage{Link: p1}
				}
				params = append(params, &param)
			}
		}

		return params, nil
	}

	return nil, nil
}
