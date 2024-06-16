package dto

import "botwhatsapp/internal/interfaces/webhooks/model"

type ChannelDTO struct {
	UserPhone string
	UserName  string
	Payload   string
	MetaID    string
	Message   model.Value
}
