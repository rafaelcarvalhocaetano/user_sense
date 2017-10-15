package dto

type InputDojoCreateMsgInteractiveDTO struct {
	Header    *HeaderInformation                   `json:"header"`
	Body      *string                              `json:"body"`
	Footer    *string                              `json:"footer"`
	MetaWbaID *string                              `json:"meta_wba_id"`
	Name      *string                              `json:"name"`
	Type      *string                              `json:"message_type,omitempty"`
	Buttons   []*DojoButtonCreateMsgInteractiveDTO `json:"buttons"`
}

type HeaderInformation struct {
	Type *string `json:"type"` //video, image, document, text
	Data *string `json:"data"`
}

type Media struct {
}

type DojoButtonCreateMsgInteractiveDTO struct {
	Id     string `json:"id,omitempty"`
	Type   string `json:"type,omitempty"`
	Title  string `json:"title"`
	NextID string `json:"next_id"`
}

type OutputDojoCreateMsgInteractiveDTO struct {
	Message string `json:"message"`
}

type InteractiveDispatchParams struct {
	Name   *string        `json:"name"`
	To     *string        `json:"to"`
	Header *string        `json:"header,omitempty"`
	Body   *string        `json:"body"`
	Footer *string        `json:"footer,omitempty"`
	Params map[string]any `json:"params"`
}
