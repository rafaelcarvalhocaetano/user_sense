package dojo

import (
	"botwhatsapp/internal/app/whatsapp/entities"
)

type GetModelInteractive struct {
	repository Repository
}

func NewGetModelInteractive(repository Repository) *GetModelInteractive {
	return &GetModelInteractive{repository: repository}
}

func (cib *GetModelInteractive) GetModelInteractive(id, name *string) (*entities.Interactive, error) {
	if id != nil {
		data, err := cib.repository.DojoGetMsgInteractiveById(*id)
		if err != nil {
			return nil, err
		}
		return cib.mapperInteractive(data), nil
	}

	if name != nil {
		data, err := cib.repository.DojoGetMsgInteractiveByName(*name)
		if err != nil {
			return nil, err
		}
		return cib.mapperInteractive(data), nil
	}

	return nil, nil
}

func (cib *GetModelInteractive) mapperInteractive(data *entities.Interactive) *entities.Interactive {
	if data.InteractiveData.Header != nil {
		header := data.InteractiveData.Header
		if header.Type == "image" {
			data.InteractiveData.Header.Image = &entities.HeaderMedia{Link: header.Data}
		}
		if header.Type == "document" {
			data.InteractiveData.Header.Document = &entities.HeaderMedia{Link: header.Data}
		}
		if header.Type == "video" {
			data.InteractiveData.Header.Video = &entities.HeaderMedia{Link: header.Data}
		}
		if header.Type == "text" {
			data.InteractiveData.Header.Text = "[here]"
		}
	}

	if data.InteractiveData.Body != nil {
		data.InteractiveData.Body.Text = "[here]"
	}

	if data.InteractiveData.Footer != nil {
		data.InteractiveData.Footer.Text = "[here]"
	}

	data.InteractiveData.Header.Data = ""
	if len(data.InteractiveData.Action.Buttons) > 0 {
		for i, _ := range data.InteractiveData.Action.Buttons {
			data.InteractiveData.Action.Buttons[i].Reply.Id = "[here]"
			data.InteractiveData.Action.Buttons[i].Reply.Title = "[here]"
		}
	} else {
		data.InteractiveData.Action = nil
	}

	return data
}
