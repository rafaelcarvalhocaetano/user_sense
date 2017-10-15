package entities

type TemplateMain struct {
	MessagingProduct string    `json:"messaging_product"`
	RecipientType    string    `json:"recipient_type"`
	To               string    `json:"to"`
	Type             string    `json:"type"`
	Template         *Template `json:"template"`
}

type Template struct {
	//ID         *string     `json:"id"`
	Name       string      `json:"name,omitempty"`
	Language   *Language   `json:"language,omitempty"`
	Components []Component `json:"components,omitempty"`
	CID        string      `json:"company_id,omitempty"`
	Category   string      `json:"category,omitempty"` //AUTHENTICATION, MARKETING, UTILITY
}

type Language struct {
	Code string `json:"code"`
}

type Component struct {
	ID      string             `json:"id,omitempty"`     // HEADER, BODY, FOOTER, BUTTONS
	Type    string             `json:"type,omitempty"`   // HEADER, BODY, FOOTER, BUTTONS
	Format  *string            `json:"format,omitempty"` // TEXT, IMAGE, LOCATION, DOCUMENT
	Text    string             `json:"text,omitempty"`
	Buttons []*ButtonComponent `json:"buttons,omitempty"`
	Params  []*Parameter       `json:"parameters,omitempty"`
	Data    []*DataParam       `json:"data,omitempty"`
}

type DataParam struct {
	Value    string `json:"value"`
	Type     string `json:"type"`
	Position string `json:"position"`
}

type ButtonComponent struct {
	Type        string       `json:"type"` // URL, PHONE_NUMBER, QUICK_REPLY, COPY_CODE
	Text        string       `json:"text"`
	PhoneNumber *string      `json:"phone_number,omitempty"`
	Url         *string      `json:"url,omitempty"`
	SubType     *string      `json:"sub_type,omitempty"`
	Index       *string      `json:"index,omitempty"`
	Data        []*DataParam `json:"data,omitempty"`
	Params      []*Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	Type        string      `json:"type"`
	Text        *string     `json:"text,omitempty"`
	Currency    *Currency   `json:"currency,omitempty"`
	DateTime    *DateTime   `json:"date_time,omitempty"`
	Payload     *string     `json:"payload,omitempty"`
	Url         *string     `json:"url,omitempty"`
	PhoneNumber *string     `json:"phone_number,omitempty"`
	SubType     *string     `json:"sub_type,omitempty"`
	Index       *string     `json:"index,omitempty"`
	Image       *Media      `json:"image,omitempty"`
	Action      *ActionData `json:"action,omitempty"`
	Document    *Media      `json:"document,omitempty"`
	Video       *Media      `json:"video,omitempty"`
}

type Media struct {
	ID       *string   `json:"id,omitempty"`
	Link     *string   `json:"link,omitempty"`
	Provider *Provider `json:"provider,omitempty"`
	Filename *string   `json:"filename,omitempty"`
}

type Provider struct {
	Name string `json:"name"`
}

type DateTime struct {
	FallbackValue string `json:"fallback_value"`
	DayOfWeek     int    `json:"day_of_week"`
	Year          int    `json:"year"`
	Month         int    `json:"month"`
	DayOfMonth    int    `json:"day_of_month"`
	Hour          int    `json:"hour"`
	Minute        int    `json:"minute"`
	Calendar      string `json:"calendar"`
}

type ActionData struct {
	ThumbnailProductRetailerID string `json:"thumbnail_product_retailer_id"`
}

type Currency struct {
	FallbackValue string `json:"fallback_value"`
	Code          string `json:"code"`
	Amount1000    int64  `json:"amount_1000"`
}
