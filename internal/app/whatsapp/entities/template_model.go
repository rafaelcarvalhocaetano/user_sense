package entities

type TemplateMetaModel struct {
	MessagingProduct string         `json:"messaging_product"`
	RecipientType    string         `json:"recipient_type"`
	To               string         `json:"to"`
	Type             string         `json:"type"`
	TemplateHeader   TemplateHeader `json:"template"`
}

type TemplateHeader struct {
	Name      string              `json:"name"`
	Language  TemplateLanguage    `json:"language"`
	Component []TemplateComponent `json:"components"`
}

type TemplateLanguage struct {
	Code string `json:"code"`
}

type TemplateComponent struct {
	Type       string               `json:"type"`
	SubType    *string              `json:"sub_type,omitempty"`
	Index      *string              `json:"index,omitempty"`
	Parameters []*TemplateParameter `json:"parameters,omitempty"`
}

type TemplateParameter struct {
	Type     string            `json:"type"`
	Format   *string           `json:"format,omitempty"`
	Text     *string           `json:"text,omitempty"`
	Currency *TemplateCurrency `json:"currency,omitempty"`
	DateTime *TemplateDate     `json:"date_time,omitempty"`
	Image    *TemplateImage    `json:"image,omitempty"`
	Document *TemplateImage    `json:"document,omitempty"`
	Payload  *string           `json:"payload,omitempty"`
	Action   *TemplateAction   `json:"action,omitempty"`
}

type TemplateAction struct {
	ThumbnailProductRetailerId string `json:"thumbnail_product_retailer_id"`
}

type TemplateImage struct {
	Link string `json:"link"`
}

type TemplateDate struct {
	FallbackValue string `json:"fallback_value"`
	DayOfWeek     int    `json:"day_of_week"`
	Year          int    `json:"year"`
	Month         int    `json:"month"`
	DayOfMonth    int    `json:"day_of_month"`
	Hour          int    `json:"hour"`
	Minute        int    `json:"minute"`
	Calendar      string `json:"calendar"`
}

type TemplateCurrency struct {
	FallbackValue string `json:"fallback_value"`
	Code          string `json:"code"`
	Amount1000    int    `json:"amount_1000"`
}
