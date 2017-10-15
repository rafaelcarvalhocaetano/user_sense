package service

//
//import (
//	"botwhatsapp/internal/app/whatsapp/entities"
//	"botwhatsapp/internal/infra/database/repository/model"
//	"strings"
//)
//
//type MapperService struct{}
//
//func NewMapperService() *MapperService {
//	return &MapperService{}
//}
//
//func (mp *MapperService) ToPersistence(data *entities.TemplateModel) *model.PersistenceTemplate {
//	templatePersistence := model.PersistenceTemplate{
//		Name:                  data.Name,
//		Language:              data.Language,
//		Category:              data.Category,
//		MetaWbaID:             data.MetaWbaID,
//		ComponentPersistences: make([]model.PersistenceComponent, 0),
//	}
//
//	if data.Component.Header != nil {
//		headerDB := model.PersistenceComponent{
//			Type:   *data.Component.Header.Type,
//			Format: data.Component.Header.Format,
//			Text:   data.Component.Header.Text,
//		}
//		mp.mapperError(data.Component.Header.Example, &headerDB)
//		templatePersistence.ComponentPersistences = append(templatePersistence.ComponentPersistences, headerDB)
//	}
//
//	bodyDB := model.PersistenceComponent{
//		Type: *data.Component.Body.Type,
//		Text: data.Component.Body.Text,
//	}
//	mp.mapperError(data.Component.Body.Example, &bodyDB)
//	templatePersistence.ComponentPersistences = append(templatePersistence.ComponentPersistences, bodyDB)
//
//	if data.Component.Footer != nil {
//		footerDB := model.PersistenceComponent{
//			Type: *data.Component.Footer.Type,
//			Text: &data.Component.Footer.Text,
//		}
//		mp.mapperError(data.Component.Footer.Example, &footerDB)
//		templatePersistence.ComponentPersistences = append(templatePersistence.ComponentPersistences, footerDB)
//	}
//
//	if data.Component.Buttons != nil {
//		for _, btn := range data.Component.Buttons.Buttons {
//			buttonsDB := model.PersistenceComponent{
//				Type:        btn.Type,
//				Text:        &btn.Text,
//				PhoneNumber: btn.PhoneNumber,
//				Url:         btn.Url,
//				OtpType:     btn.OtpType,
//			}
//			mp.mapperError(data.Component.Buttons.Example, &buttonsDB)
//			templatePersistence.ComponentPersistences = append(templatePersistence.ComponentPersistences, buttonsDB)
//		}
//	}
//
//	return &templatePersistence
//}
//
//func (mp *MapperService) mapperError(ex *entities.Exemplo, data *model.PersistenceComponent) {
//	if ex != nil {
//		if ex.HeaderText != nil {
//			txt := strings.Join(ex.HeaderText, ",")
//			data.ExampleText = &txt
//		}
//		if ex.HeaderHandle != nil {
//			hx := strings.Join(ex.HeaderHandle, ",")
//			data.ExampleHandle = &hx
//		}
//		if ex.HeaderUrl != nil {
//			rl := strings.Join(ex.HeaderUrl, ",")
//			data.ExampleUrl = &rl
//		}
//		if ex.BodyText != nil {
//			for _, bt := range ex.BodyText {
//				btx := strings.Join(bt, ",")
//				data.BodyText = &btx
//			}
//		}
//	}
//}
