package dojo

import (
	"botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"
	"errors"
	"fmt"
	"os"
)

type DojoCreateFlow struct {
	repository  Repository
	template    *GetModelTemplateWrite
	interactive *GetModelInteractive
}

func NewDojoCreateFlow(
	repository Repository,
	temp *GetModelTemplateWrite,
	inter *GetModelInteractive) *DojoCreateFlow {
	return &DojoCreateFlow{repository: repository, template: temp, interactive: inter}
}

func (dc *DojoCreateFlow) DojoCreateFlow(input *dto.InputCreateDojoFlow) (any, error) {
	if input.Name == nil || input.Default == nil || input.Flows == nil || len(input.Flows) == 0 {
		return nil, errors.New("invalid params")
	}

	if input.Default != nil && *input.Type != "template" {
		return nil, errors.New("template default is required like template")
	}

	responseFlow, err := dc.template.GetModelTemplateByWrite(*input.Default)
	if err != nil {
		return nil, errors.New("template default not found")
	}
	if responseFlow == nil {
		return nil, errors.New("template default not found")
	}

	for i, flow := range input.Flows {
		if flow.Current == nil {
			return nil, errors.New(fmt.Sprintf("current is required in position: %v", i))
		}
		if flow.ResponseSuccess != nil && flow.ResponseError == nil {
			return nil, errors.New(fmt.Sprintf("response error: %v", flow.ResponseError))
		}
		if *flow.Type != "template" && *flow.Type != "interactive" {
			return nil, errors.New(fmt.Sprintf("type must be template or interactive"))
		}

		if *flow.Type == "template" {
			template, err := dc.template.GetModelTemplateByWrite(*flow.Current)
			if err != nil || template == nil {
				return nil, errors.New(fmt.Sprintf("template [ %v ] not found", *flow.Current))
			}
		}

		if *flow.Type == "interactive" {
			interactive, err := dc.interactive.GetModelInteractive(nil, flow.Current)
			if err != nil || interactive == nil {
				return nil, errors.New(fmt.Sprintf("interactive [ %v ] not found", *flow.Current))
			}
		}
	}

	metaID := os.Getenv("META_WBA_ID")
	if err := dc.repository.CreateFlow(metaID, input); err != nil {
		return nil, errors.New(fmt.Sprintf("create dojo error: %v", err.Error()))
	}

	return responseFlow, nil
}
