package dojo

import "botwhatsapp/internal/app/whatsapp/usecases/dojo/dto"

type DojoGetFlow struct {
	repository Repository
}

func NewDojoGetFlow(repository Repository) *DojoGetFlow {
	return &DojoGetFlow{repository: repository}
}

func (dc *DojoGetFlow) DojoGetFlow(id, name string) (*dto.FlowData, error) {
	response, err := dc.repository.GetFlowByMetaIdAndName(id, name)
	if err != nil {
		return nil, err
	}
	flow := dto.FlowData{Actions: make([]*dto.FlowAction, 0)}

	for _, v := range *response {
		flow.Name = *v.Name
		if v.Default != "" {
			flow.Default = v.Default
		}

		v.Default = ""

		flow.Actions = append(flow.Actions, &dto.FlowAction{
			Used:     false,
			Order:    v.Order,
			Current:  v.Current,
			Default:  v.Default,
			Resposta: v.Resposta,
			Error:    v.Error,
			Type:     v.Type,
		})

	}

	return &flow, nil
}
