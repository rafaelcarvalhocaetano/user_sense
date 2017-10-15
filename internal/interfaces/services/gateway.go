package services

type SendMessagesHttp interface {
	Send(p string, data any) (map[string]any, error)
}

type RegisterHttp interface {
	Register(path string, data any) (map[string]any, error)
}

type GetTemplateGateway interface {
	GetDataModel(p string, data map[string]string) (map[string]any, error)
}

type DeleteTemplateGateway interface {
	DeleteTemplate(name string) error
}

type WTAGateway struct {
	SendMessagesHttp
	RegisterHttp
	GetTemplateGateway
	DeleteTemplateGateway
}
