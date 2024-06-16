package channel

import (
	"sync"

	"botwhatsapp/internal/interfaces/webhooks/model"
)

type MessageChannel struct{}

func NewMessageChannel() *MessageChannel {
	return &MessageChannel{}
}

func (cm *MessageChannel) Main(webhook <-chan *model.WebhookData, send chan<- *UserData, templateStatus chan<- model.Status) {
	var wg sync.WaitGroup

	go func() {
		for c := range webhook {
			if len(c.Entry) == 0 {
				continue
			}

			entry := c.Entry[0]
			if len(entry.Changes) == 0 {
				continue
			}

			change := entry.Changes[0]
			phone := cm.userPhoneNumber(*change)
			dataChannel := &UserData{UserPhone: *phone, MetaID: entry.ID}

			wg.Add(1)
			go func() {
				defer wg.Done()

				statuses := change.Value.Statuses
				contacts := change.Value.Contacts
				messages := change.Value.Messages

				if len(contacts) > 0 {
					profile := contacts[0].Profile
					dataChannel.UserName = profile.Name
				}

				if len(statuses) > 0 {
					// TODO: bloco de codigo sobre o status dos templates
					templateStatus <- *statuses[0]
					return
				}

				if len(messages) > 0 {
					message := change.Value.Messages[0]
					dataChannel.Message = *change.Value
					switch message.Type {
					case "text":
						dataChannel.Payload = message.Text.Body
					case "button":
						if message.Button.Payload != nil {
							dataChannel.Payload = *message.Button.Payload
						}
					case "interactive":
						interactive := message.Interactive
						if interactive != nil && interactive.Type == "button_reply" {
							dataChannel.Payload = interactive.ButtonReply.ID
						}
					}
				}
				send <- dataChannel
			}()
		}
		wg.Wait()
		//close(send)
	}()
}

func (cm *MessageChannel) userPhoneNumber(webmessage model.Change) *string {
	var userPhone string

	if statuses := webmessage.Value.Statuses; statuses != nil && len(statuses) > 0 {
		if recipientID := statuses[0].RecipientID; recipientID != nil {
			userPhone = *recipientID
		}
	}

	if userPhone == "" {
		if contacts := webmessage.Value.Contacts; contacts != nil && len(contacts) > 0 {
			if waID := contacts[0].WaID; waID != nil {
				userPhone = *waID
			}
		}
	}

	return &userPhone
}
