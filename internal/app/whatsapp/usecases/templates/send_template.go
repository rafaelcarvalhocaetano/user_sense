package templates

//
//import (
//	"botwhatsapp/internal/app/whatsapp"
//	"botwhatsapp/internal/interfaces/services"
//)
//
//type SendTemplate struct {
//	gateway services.WTAGateway
//	db      whatsapp.Repository
//}
//
//func NewSendTemplate(db whatsapp.Repository, gateway services.WTAGateway) *SendTemplate {
//	return &SendTemplate{db: db, gateway: gateway}
//}
//
//func (send *SendTemplate) SendTemplate(data *map[string]any) (*map[string]any, error) {
//	response := make(map[string]any)
//	resp, err := send.gateway.Send("messages", data)
//	if err != nil {
//		response["error"] = err.Error()
//		return nil, err
//	}
//
//	response["success"] = resp
//	return &response, nil
//}
