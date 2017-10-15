package entities

type Interactive struct {
	InteractiveData  *InteractiveData `json:"interactive"`
	MessagingProduct string           `json:"messaging_product"`
	RecipientType    string           `json:"recipient_type"`
	Type             string           `json:"type"`
	To               string           `json:"to"`
}

type InteractiveData struct {
	Name   string    `json:"name,omitempty"`
	Type   *string   `json:"type,omitempty"`
	Header *Header   `json:"header,omitempty"`
	Body   *ItemData `json:"body,omitempty"`
	Footer *ItemData `json:"footer,omitempty"`
	Action *Action   `json:"action,omitempty"`
}

type Action struct {
	Buttons []*Buttons `json:"buttons,omitempty"`
	//Button   string     `json:"button"`
	//Sections []*Section `json:"sections"`
}

type Header struct {
	Type     string       `json:"type"`
	Data     string       `json:"data,omitempty"`
	Text     string       `json:"text,omitempty"`
	Image    *HeaderMedia `json:"image,omitempty"`
	Document *HeaderMedia `json:"document,omitempty"`
	Video    *HeaderMedia `json:"video,omitempty"`
}

type HeaderMedia struct {
	Link string `json:"link"`
}

type ItemData struct {
	Text string `json:"text"`
}

//type Section struct {
//	Title string `json:"title"`
//	Rows  []*Row `json:"rows"`
//}

//type Row struct {
//	Id          string `json:"id"`
//	Title       string `json:"title"`
//	Description string `json:"description"`
//}

type Buttons struct {
	Type  string `json:"type"`
	Reply *Reply `json:"reply"`
}

type Reply struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
