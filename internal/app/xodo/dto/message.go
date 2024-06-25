package dto

type OutputMessage struct {
	Message string `json:"body"`
}

type InputMessage struct {
	To            string `json:"to"`
	Message       string `json:"message"`
	ContactNumber string `json:"contact_number,omitempty"`
	ContactName   string `json:"contact_name,omitempty"`
	Type          string `json:"type"`
	Caption       string `json:"caption,omitempty"`
	Link          string `json:"link,omitempty"`
	MessageID     string `json:"message_id,omitempty"`
}
