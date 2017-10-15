package whatsapp

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	wtaMessage "botwhatsapp/internal/app/whatsapp/usecases/message/dto"
	_temp "botwhatsapp/internal/app/whatsapp/usecases/templates/dto"
)

type RegisterGateway interface {
	ExecuteRegister(data *entities.TemplateModel) (map[string]any, error)
}

type DispatchTemplateGateway interface {
	ExecuteDispatch(input *_temp.Dispatch) ([]map[string]any, error)
}

type GetTemplateMetaGateway interface {
	GetTemplateMeta(input map[string]string) (*[]entities.TemplateMetaModel, error)
}

type GetTemplateMetaWithPhoneGateway interface {
	GetTemplateMetaWithPhone(input map[string]string) ([]entities.TemplateMetaModel, error)
}

type SendTemplateGateway interface {
	SendTemplate(data *map[string]any) (*map[string]any, error)
}

type SimpleMessageGatway interface {
	SimpleMessage(data *wtaMessage.MessageData) (*wtaMessage.OutputSimpleMessage, error)
}

type Gateway struct {
	RegisterGateway
	DispatchTemplateGateway
	GetTemplateMetaGateway
	SendTemplateGateway
	SimpleMessageGatway
	GetTemplateMetaWithPhoneGateway
}
