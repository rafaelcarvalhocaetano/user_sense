package dojo

import (
	"botwhatsapp/internal/interfaces/services"
)

type MainDOJO struct{}

func New() *MainDOJO {
	return &MainDOJO{}
}
func (*MainDOJO) Main(dojoRepo Repository, http services.WTAGateway) *Gateway {
	dojoCreateMsgInteractive := NewDojoCreateMsgInteractive(dojoRepo)
	dojoGetMsgInteractive := NewGetModelInteractive(dojoRepo)
	getModelTemplateWrite := NewGetModelTemplateWrite(http)
	registerTemplateMeta := NewExecuteRegister(dojoRepo, http)
	saveStatusUserMessage := NewSaveStatusUserMessage(dojoRepo)
	sendOneTemplate := NewSendOneTemplate(http, *saveStatusUserMessage)
	disaptcher := NewDispatcher(dojoRepo, *getModelTemplateWrite, http, *saveStatusUserMessage)
	sendSimpleMessage := NewSendSimpleMessage(http)
	modelRegisterTemplate := NewModelRegisterTemplate()
	createFlow := NewDojoCreateFlow(dojoRepo, getModelTemplateWrite, dojoGetMsgInteractive)
	getFlow := NewDojoGetFlow(dojoRepo)
	dispatchInteractive := NewDispatchInteractive(dojoRepo, http)

	gt := Gateway{
		CreateMsgInteractiveGateway:  dojoCreateMsgInteractive,
		GetMsgInteractiveGateway:     dojoGetMsgInteractive,
		GetModelTemplateWriteGateway: getModelTemplateWrite,
		RegisterTemplateMetaGateway:  registerTemplateMeta,
		SendOneTemplateGateway:       sendOneTemplate,
		DispatcherGateway:            disaptcher,
		SaveStatusUserMessageGateway: saveStatusUserMessage,
		SendSimpleMessageGateway:     sendSimpleMessage,
		ModelRegisterTemplateGateway: modelRegisterTemplate,
		DojoCreateFlowGateway:        createFlow,
		DojoGetFlowGateway:           getFlow,
		DispatchInteractiveGateway:   dispatchInteractive,
	}

	return &gt
}
