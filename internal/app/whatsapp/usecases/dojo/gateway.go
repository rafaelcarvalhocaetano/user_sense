package dojo

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"
)

type Repository interface {
	DojoCreateMsgInteractive(data *dto.InputDojoCreateMsgInteractiveDTO) error
	DojoGetMsgInteractiveById(id string) (*entities.Interactive, error)
	DojoGetMsgInteractiveByName(name string) (*entities.Interactive, error)
	RegisterTemplate(name, category, mid *string) error
	DeleteTemplate(name string) error
	DojoSaveUserMessage(name, phone, message, status, sID, recipient *string) error
	CreateFlow(mid string, data *dto.InputCreateDojoFlow) error
	GetFlowByMetaIdAndName(mid, name string) (*[]dto.FlowAction, error)
}

type CreateMsgInteractiveGateway interface {
	DojoCreateMsgInteractive(input *dto.InputDojoCreateMsgInteractiveDTO) (
		*dto.OutputDojoCreateMsgInteractiveDTO, error,
	)
}

type GetMsgInteractiveGateway interface {
	GetModelInteractive(id, name *string) (*entities.Interactive, error)
}

type GetModelTemplateWriteGateway interface {
	GetModelTemplateByWrite(name string) ([]entities.TemplateMetaModel, error)
}

type RegisterTemplateMetaGateway interface {
	RegisterTemplateMeta(data *entities.TemplateModel) (map[string]any, error)
}

type SendOneTemplateGateway interface {
	SendOneTemplate(data map[string]any) (*map[string]any, error)
}

type DispatcherGateway interface {
	Dispatch(input *dto.DojoDispatchParams) (map[string]interface{}, error)
}

type SaveStatusUserMessageGateway interface {
	SaveStatusUserMessage(to string, response map[string]interface{}) error
}

type SendSimpleMessageGateway interface {
	SendSimpleMessage(data *InputSendSimpleMessage) (*outputSendSimpleMessage, error)
}

type ModelRegisterTemplateGateway interface {
	ModelRegisterTemplate() any
}

type DojoCreateFlowGateway interface {
	DojoCreateFlow(input *dto.InputCreateDojoFlow) (any, error)
}

type DojoGetFlowGateway interface {
	DojoGetFlow(id, name string) (*dto.FlowData, error)
}

type DispatchInteractiveGateway interface {
	DispatchInteractive(input *dto.InteractiveDispatchParams) (any, error)
}

type Gateway struct {
	CreateMsgInteractiveGateway
	GetMsgInteractiveGateway
	GetModelTemplateWriteGateway
	RegisterTemplateMetaGateway
	SendOneTemplateGateway
	DispatcherGateway
	SaveStatusUserMessageGateway
	SendSimpleMessageGateway
	ModelRegisterTemplateGateway
	DojoCreateFlowGateway
	DojoGetFlowGateway
	DispatchInteractiveGateway
}
