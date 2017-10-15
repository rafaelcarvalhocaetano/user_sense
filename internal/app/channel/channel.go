package channel

import "botwhatsapp/internal/interfaces/webhooks/model"

type UserData struct {
	UserPhone string
	UserName  string
	Payload   string
	MetaID    string
	Message   model.Value
}
