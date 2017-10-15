package templates

//
//import (
//	"botwhatsapp/internal/app/whatsapp"
//	"botwhatsapp/internal/app/whatsapp/entities"
//	"botwhatsapp/internal/app/whatsapp/service"
//	"botwhatsapp/internal/interfaces/services"
//)
//
//type ExecuteRegister struct {
//	http services.WTAGateway
//}
//
//type OutputRegisterModel struct {
//	Message string `json:"message"`
//}
//
//func NewExecuteRegister(db whatsapp.Repository, http services.WTAGateway) *ExecuteRegister {
//	return &ExecuteRegister{db: db, http: http}
//}
//
//func (temp *ExecuteRegister) ExecuteRegister(data *entities.TemplateModel) (map[string]any, error) {
//	tempData, err := entities.NewTemplateModel(data)
//	if err != nil {
//		return nil, err
//	}
//
//	if data.Component.Header != nil {
//		header, errHeader := entities.NewHeaderModel(data.Component.Header)
//		if errHeader != nil {
//			return nil, errHeader
//		}
//		tempData.Component.Header = header
//	}
//
//	body, err := entities.NewBodyModel(data.Component.Body)
//	if err != nil {
//		return nil, err
//	}
//	tempData.Component.Body = body
//
//	if data.Component.Footer != nil {
//		footer, errF := entities.NewFooterModel(data.Component.Footer)
//		if errF != nil {
//			return nil, errF
//		}
//		tempData.Component.Footer = footer
//	}
//
//	if data.Component.Buttons != nil {
//		buttons, errBtn := entities.NewButtonModel(data.Component.Buttons)
//		if errBtn != nil {
//			return nil, errBtn
//		}
//		tempData.Component.Buttons = buttons
//	}
//
//	mapperService := service.NewMapperService()
//	mappedTemplate := mapperService.ToPersistence(tempData)
//	response := make(map[string]any)
//	response["data"] = tempData
//	response["persistence"] = mappedTemplate
//	response["request"] = tempData.RequestData()
//
//	//respHttp, err := temp.http.Register("message_templates", tempData.RequestData())
//	//if err != nil {
//	//	return nil, err
//	//}
//	//if respHttp != nil {
//	//	err = temp.db.UploadTemplate(mappedTemplate)
//	//	if err != nil {
//	//		_ = temp.http.DeleteTemplate(mappedTemplate.Name)
//	//		return nil, err
//	//	}
//	//}
//	err = temp.db.UploadTemplate(mappedTemplate)
//	if err != nil {
//		_ = temp.http.DeleteTemplate(mappedTemplate.Name)
//		return nil, err
//	}
//
//	return response, nil
//}
