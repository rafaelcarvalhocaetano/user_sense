package dojo

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"
	"botwhatsapp/internal/interfaces/services"
	"errors"
	"fmt"
	"strings"
)

type Dispatcher struct {
	repository Repository
	tmodel     GetModelTemplateWrite
	gateway    services.WTAGateway
	saveStatus SaveStatusUserMessage
}

func NewDispatcher(
	repository Repository,
	tmodel GetModelTemplateWrite,
	gateway services.WTAGateway,
	saveStatus SaveStatusUserMessage,
) *Dispatcher {
	return &Dispatcher{
		gateway:    gateway,
		repository: repository,
		tmodel:     tmodel,
		saveStatus: saveStatus,
	}
}

func (d *Dispatcher) Dispatch(input *dto.DojoDispatchParams) (map[string]interface{}, error) {
	if input.Name == nil || input.To == nil {
		return nil, errors.New("id in dispatcher is required")
	}

	template, err := d.tmodel.GetModelTemplateByWrite(*input.Name)
	if err != nil {
		return nil, err
	}
	if template == nil {
		return nil, nil
	}

	template[0].To = *input.To
	if input.Params != nil {
		replaceText(&template[0], *input.Params)
	}
	temps := map[string]interface{}{
		"template": template[0],
	}
	resp, err := d.gateway.Send("messages", template[0])
	if err != nil {
		temps["error"] = err.Error()
		return nil, err
	}

	err = d.saveStatus.SaveStatusUserMessage(template[0].To, resp)
	if err != nil {
		temps["error"] = err.Error()
	}

	temps["resp"] = resp
	return temps, nil
}

func replaceText(template *entities.TemplateMetaModel, params map[string]interface{}) {
	for _, component := range template.TemplateHeader.Component {
		componentType := strings.ToLower(component.Type)
		for i, parameter := range component.Parameters {
			if strings.ToUpper(parameter.Type) == "TEXT" && parameter.Text != nil && *parameter.Text == "[here]" {
				var placeholderKey string
				if componentType == "button" {
					placeholderKey = fmt.Sprintf("param_next_%d", i+1)
				} else {
					placeholderKey = fmt.Sprintf("param_%s_%d", componentType, i+1)
				}
				if paramValue, exists := params[placeholderKey]; exists {
					paramStr := paramValue.(string)
					parameter.Text = &paramStr
				}
			}
			if strings.ToUpper(parameter.Type) == "IMAGE" && parameter.Image != nil && parameter.Image.Link == "[here]" {
				var placeholderKey string
				placeholderKey = fmt.Sprintf("param_%s_%d", componentType, i+1)
				if paramValue, exists := params[placeholderKey]; exists {
					paramStr := paramValue.(string)
					parameter.Image.Link = paramStr
				}
			}
		}
	}
}
