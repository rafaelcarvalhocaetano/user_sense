package dojo

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"
	"botwhatsapp/internal/interfaces/services"
	"errors"
	"fmt"
	"regexp"
)

type DispatchInteractive struct {
	repository Repository
	gateway    services.WTAGateway
}

func NewDispatchInteractive(repository Repository, gateway services.WTAGateway) *DispatchInteractive {
	return &DispatchInteractive{gateway: gateway, repository: repository}
}

func (ind *DispatchInteractive) DispatchInteractive(input *dto.InteractiveDispatchParams) (any, error) {
	if input.Name == nil || input.To == nil {
		return nil, errors.New("parameters is required")
	}

	data, err := ind.repository.DojoGetMsgInteractiveByName(*input.Name)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("interactive message not found")
	}

	if input.Header != nil {
		header := data.InteractiveData.Header
		if header.Type == "image" || header.Type == "document" || header.Type == "video" {
			urlRegex := `^(https?):\/\/[^\s/$.?#].[^\s]*$`
			regex := regexp.MustCompile(urlRegex)
			if !regex.MatchString(*input.Header) {
				return nil, errors.New("header is required url")
			}
		}

		if header.Type == "image" {
			data.InteractiveData.Header.Image = &entities.HeaderMedia{Link: *input.Header}
		}
		if header.Type == "document" {
			data.InteractiveData.Header.Document = &entities.HeaderMedia{Link: *input.Header}
		}
		if header.Type == "video" {
			data.InteractiveData.Header.Video = &entities.HeaderMedia{Link: *input.Header}
		}
		if header.Type == "text" {
			data.InteractiveData.Header.Text = *input.Header
		}
		data.InteractiveData.Header.Data = ""
	} else {
		data.InteractiveData.Header = nil
	}

	if input.Body != nil {
		data.InteractiveData.Body.Text = *input.Body
	}

	fmt.Println("input: ", input.Footer)
	if input.Footer != nil {
		data.InteractiveData.Footer.Text = *input.Footer
	}

	if data.InteractiveData.Footer.Text == "" {
		data.InteractiveData.Footer = nil
	}

	for i, _ := range data.InteractiveData.Action.Buttons {
		keyID := fmt.Sprintf("button_%v_id", i+1)
		id, ok := input.Params[keyID]
		if !ok {
			return nil, errors.New(fmt.Sprintf("invalid parameters id %v", keyID))
		}

		keyTitle := fmt.Sprintf("button_%v_title", i+1)
		title, ok := input.Params[keyTitle]
		if !ok {
			return nil, errors.New(fmt.Sprintf("invalid parameters title %v", keyTitle))
		}

		data.InteractiveData.Action.Buttons[i].Reply.Id = id.(string)
		data.InteractiveData.Action.Buttons[i].Reply.Title = title.(string)
	}

	data.To = *input.To
	resp, err := ind.gateway.Send("messages", data)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
