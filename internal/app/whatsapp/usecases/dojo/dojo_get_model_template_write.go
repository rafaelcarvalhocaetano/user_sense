package dojo

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"botwhatsapp/internal/app/whatsapp/service"
	"botwhatsapp/internal/interfaces/services"
	"errors"
	"strings"
)

type GetModelTemplateWrite struct {
	http services.WTAGateway
}

func NewGetModelTemplateWrite(http services.WTAGateway) *GetModelTemplateWrite {
	return &GetModelTemplateWrite{http: http}
}

func (temp *GetModelTemplateWrite) GetModelTemplateByWrite(name string) ([]entities.TemplateMetaModel, error) {
	mapName := map[string]string{"name": name}
	resp, err := temp.http.GetDataModel("message_templates", mapName)
	if err != nil {
		return nil, err
	}

	model, err := temp.execute("55(xx)xxxxxxxxx", resp)
	if err != nil {
		return nil, err
	}

	return *model, nil
}

func (temp *GetModelTemplateWrite) execute(phone string, resp map[string]any) (*[]entities.TemplateMetaModel, error) {
	templates, ok := resp["data"].([]any)
	if !ok {
		return nil, errors.New("invalid response")
	}

	var templateResponse []entities.TemplateMetaModel
	for _, template := range templates {
		tmp := template.(map[string]any)
		templateData := entities.TemplateMetaModel{
			MessagingProduct: "whatsapp",
			RecipientType:    "individual",
			To:               phone,
			Type:             "template",
			TemplateHeader: entities.TemplateHeader{
				Name:      tmp["name"].(string),
				Component: make([]entities.TemplateComponent, 0),
				Language:  entities.TemplateLanguage{Code: "pt_BR"},
			},
		}

		// Start Validate Params
		components := tmp["components"].([]any)
		for _, c := range components {
			component := c.(map[string]any)
			componentType := strings.ToUpper(component["type"].(string))
			typeComponentFormated := entities.ComponentType(componentType)
			templateComponent := entities.TemplateComponent{Parameters: make([]*entities.TemplateParameter, 0)}

			// HEADER
			if typeComponentFormated == entities.ComponentTypeHeader {
				headerParams := service.NewHeaderValidate()
				params, err := headerParams.CheckParameters(component)
				if err != nil {
					return nil, err
				}

				if params != nil {
					templateComponent.Type = componentType
					templateComponent.Parameters = append(templateComponent.Parameters, params...)
					templateData.TemplateHeader.Component = append(templateData.TemplateHeader.Component, templateComponent)
				}
			}

			// BODY
			if typeComponentFormated == entities.ComponentTypeBody {
				bodyValidate := service.NewBodyValidate()
				params, err := bodyValidate.CheckParameters(component)
				if err != nil {
					return nil, err
				}

				if params != nil {
					templateComponent.Type = componentType
					templateComponent.Parameters = append(templateComponent.Parameters, params...)
					templateData.TemplateHeader.Component = append(templateData.TemplateHeader.Component, templateComponent)
				}
			}

			// FOOTER
			if typeComponentFormated == entities.ComponentTypeFooter {
				footerData := service.NewFooterValidate()
				footerCheck, err := footerData.CheckParameters(component)
				if err != nil {
					return nil, err
				}
				if footerCheck != nil {
					templateComponent.Type = componentType
					templateComponent.Parameters = append(templateComponent.Parameters, footerCheck...)
					templateData.TemplateHeader.Component = append(templateData.TemplateHeader.Component, templateComponent)
				}
			}

			// BUTTONS
			if typeComponentFormated == entities.ComponentTypeButtons {
				buttons := service.NewButtonsValidate()
				buttonsData, err := buttons.CheckParameters(component)
				if err != nil {
					return nil, err
				}

				if buttonsData != nil {
					templateComponent.Type = componentType
					//templateComponent.Parameters = append(templateComponent.Parameters, buttonsData...)
					templateData.TemplateHeader.Component = append(templateData.TemplateHeader.Component, *buttonsData...)
				}
			}
		}
		templateResponse = append(templateResponse, templateData)
	}

	return &templateResponse, nil
}
