package entities

type DataModel struct {
	Data Data `json:"data,omitempty"`
}

type Data struct {
	Object string      `json:"object,omitempty"`
	Entry  []EntryData `json:"entry,omitempty"`
}

type EntryData struct {
	Id      string        `json:"id,omitempty"`
	Changes []ChangesData `json:"changes,omitempty"`
}

type ChangesData struct {
	Value    *ValueData     `json:"value,omitempty"`
	Field    *string        `json:"field,omitempty"`
	Messages *[]MessageItem `json:"messages,omitempty"`
	Statuses *StatusesData  `json:"statuses,omitempty"`
}

type ValueData struct {
	MessagingProduct *string        `json:"messaging_product,omitempty"`
	Metadata         *Metadata      `json:"metadata,omitempty"`
	Contacts         *[]ContactData `json:"contacts,omitempty"`
}

type Metadata struct {
	DisplayPhoneNumber *string `json:"display_phone_number,omitempty"`
	PhoneNumberId      *string `json:"phone_number_id,omitempty"`
}

type ContactData struct {
	Profile *ProfileData `json:"profile,omitempty"`
	WaId    *string      `json:"wa_id,omitempty"`
}

type ProfileData struct {
	Name *string `json:"name,omitempty"`
}

type MessageItem struct {
	From      *string   `json:"from,omitempty"`
	Id        *string   `json:"id,omitempty"`
	Timestamp *string   `json:"timestamp,omitempty"`
	Type      *string   `json:"type,omitempty"`
	Text      *TextData `json:"text,omitempty"`
}

type TextData struct {
	Body *string `json:"body,omitempty"`
}

type StatusesData struct {
	Id           *string           `json:"id,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Timestamp    *string           `json:"timestamp,omitempty"`
	RecipientId  *string           `json:"recipient_id,omitempty"`
	Type         *string           `json:"type,omitempty"`
	Conversation *ConversationData `json:"conversation,omitempty"`
	Pricing      *PricingData      `json:"pricing,omitempty"`
}

type PricingData struct {
	PricingModel *string `json:"pricing_model,omitempty"`
	Billing      *bool   `json:"billing,omitempty"`
	Category     *string `json:"category,omitempty"`
}

type ConversationData struct {
	Id     *string     `json:"id,omitempty"`
	Origin *OriginData `json:"origin,omitempty"`
}

type OriginData struct {
	Type *string `json:"type,omitempty"`
}

//var respClient any
//v := input.Entry[0].Changes[0].Value
//if v.Messages != nil {
//	msg := v.Messages[0]
//	var message string
//	if msg.Text != nil {
//		message = msg.Text.Body
//	}
//	interactive := msg.Interactive
//	from := msg.From
//
//	// regra para o disparo do primeiro template ou menu inicial
//	if message != "" {
//		resp, errInter := wt.interactive.FindInteractiveByID("00d7448b-f1ec-41e1-9ec5-cf6e1a9865e7")
//		if errInter != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			_, _ = w.Write(wt.output.Error(input, errInter.Error()))
//		}
//
//		_ = model.MessageDispatch{
//			MessagingProduct: "whatsapp",
//			RecipientType:    "individual",
//			To:               from,
//			Type:             "interactive",
//			Interactive:      resp,
//		}
//		//respClient, errInter = wt.httpClinet.Send("messages", mms)
//		//if errInter != nil {
//		//	w.WriteHeader(http.StatusBadRequest)
//		//	_, _ = w.Write(wt.output.Error(input, errInter.Error()))
//		//}
//		return
//	}
//
//	if interactive != nil {
//		var xid string
//		buttonReply := interactive.ButtonReply
//		listReply := interactive.ListReply
//
//		listItems := make([]string, 0)
//
//		if buttonReply != nil {
//			listItems = append(listItems, buttonReply.Title)
//			xid = listReply.ID
//		}
//
//		if listItems != nil {
//			listItems = append(listItems, listReply.Title)
//			xid = listReply.ID
//		}
//
//		resp, errInter := wt.interactive.FindInteractiveByID(xid)
//		if errInter != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			_, _ = w.Write(wt.output.Error(input, errInter.Error()))
//		}
//
//		mms := model.MessageDispatch{
//			MessagingProduct: "whatsapp",
//			RecipientType:    "individual",
//			To:               from,
//			Type:             "interactive",
//		}
//		if *resp.Type == "text" {
//			mms.Type = "text"
//			mms.Text = &model.Text{
//				Body: fmt.Sprintf("Os produtos %v estão sendo processado ...", strings.Join(listItems, ", ")),
//			}
//		} else {
//			mms.Interactive = resp
//		}
//
//		//respClient, errInter = wt.httpClinet.Send("messages", mms)
//		//if errInter != nil {
//		//	w.WriteHeader(http.StatusBadRequest)
//		//	_, _ = w.Write(wt.output.Error(input, errInter.Error()))
//		//}
//		return
//	}
//}
