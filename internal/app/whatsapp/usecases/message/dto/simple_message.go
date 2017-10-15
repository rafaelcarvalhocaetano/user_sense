package dto

type MessageData struct {
	MessagingProduct string `json:"messaging_product,omitempty"`
	To               string `json:"to"`
	Text             Text   `json:"text"`
}

type Text struct {
	Body string `json:"body"`
}

type OutputSimpleMessage struct {
	Message string `json:"body"`
}
