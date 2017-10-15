package dojo

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"botwhatsapp/internal/interfaces/services"
)

type ExecuteRegister struct {
	db   Repository
	http services.WTAGateway
}

type OutputRegisterModel struct {
	Message string `json:"message"`
}

func NewExecuteRegister(db Repository, http services.WTAGateway) *ExecuteRegister {
	return &ExecuteRegister{db: db, http: http}
}

func (temp *ExecuteRegister) RegisterTemplateMeta(data *entities.TemplateModel) (map[string]any, error) {
	tempData, err := entities.NewTemplateModel(data)
	if err != nil {
		return nil, err
	}

	if data.Component.Header != nil {
		header, errHeader := entities.NewHeaderModel(data.Component.Header)
		if errHeader != nil {
			return nil, errHeader
		}
		tempData.Component.Header = header
	}

	body, err := entities.NewBodyModel(data.Component.Body)
	if err != nil {
		return nil, err
	}
	tempData.Component.Body = body

	if data.Component.Footer != nil {
		footer, errF := entities.NewFooterModel(data.Component.Footer)
		if errF != nil {
			return nil, errF
		}
		tempData.Component.Footer = footer
	}

	if data.Component.Buttons != nil {
		buttons, errBtn := entities.NewButtonModel(data.Component.Buttons)
		if errBtn != nil {
			return nil, errBtn
		}
		tempData.Component.Buttons = buttons
	}

	response := make(map[string]any)
	response["data"] = tempData
	response["request"] = tempData.RequestData()

	//mid := os.Getenv("META_WBA_ID")
	//err = temp.db.RegisterTemplate(&tempData.Name, &tempData.Category, &mid)
	//if err != nil {
	//	return nil, err
	//}
	respMeta, err := temp.http.Register("message_templates", tempData.RequestData())
	if err != nil {
		//_ = temp.db.DeleteTemplate(tempData.Name)
		return nil, err
	}
	response["meta"] = respMeta

	return response, nil
}
