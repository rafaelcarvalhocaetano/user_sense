package dojo

import (
	"botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"
	"errors"
	"regexp"
)

type CreateMsgInteractive struct {
	repository Repository
}

func NewDojoCreateMsgInteractive(db Repository) *CreateMsgInteractive {
	return &CreateMsgInteractive{repository: db}
}

func (cib *CreateMsgInteractive) DojoCreateMsgInteractive(input *dto.InputDojoCreateMsgInteractiveDTO) (*dto.
	OutputDojoCreateMsgInteractiveDTO, error) {
	if input.MetaWbaID == nil {
		return nil, errors.New("invalid department ID")
	}
	if input.Body == nil {
		return nil, errors.New("body is required")
	}

	if input.Header != nil {
		if input.Header.Type == nil || input.Header.Data == nil {
			return nil, errors.New("header params is required")
		}
		if *input.Header.Type == "image" || *input.Header.Type == "document" || *input.Header.Type == "video" {
			urlRegex := `^(https?):\/\/[^\s/$.?#].[^\s]*$`
			regex := regexp.MustCompile(urlRegex)
			if !regex.MatchString(*input.Header.Data) {
				return nil, errors.New("invalid url")
			}
		}
	}

	if len(input.Buttons) < 1 && len(input.Buttons) > 3 {
		return nil, errors.New("buttons length between 1 and 3")
	}

	for _, i := range input.Buttons {
		i.Type = "reply"
	}
	t := "buttons"
	input.Type = &t

	if err := cib.repository.DojoCreateMsgInteractive(input); err != nil {
		return nil, err
	}

	return &dto.OutputDojoCreateMsgInteractiveDTO{Message: "Save dojo message interactive"}, nil
}
