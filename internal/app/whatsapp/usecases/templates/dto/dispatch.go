package dto

type Dispatch struct {
	TemplateName string   `json:"template_name"`
	Clients      []Client `json:"clients"`
}

type Client struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Value    string `json:"value"`
	Datetime string `json:"datetime"`
}
type Message struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Template         Template `json:"template"`
}

type Template struct {
	Name       string      `json:"name"`
	Language   Language    `json:"language"`
	Components []Component `json:"components"`
}

type Language struct {
	Code string `json:"code"`
}

type Component struct {
	Type       string      `json:"type"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	Currency *Currency `json:"currency,omitempty"`
	DateTime *DateTime `json:"date_time,omitempty"`
}

type Currency struct {
	FallbackValue string `json:"fallback_value"`
	Code          string `json:"code"`
	Amount1000    int    `json:"amount_1000"`
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
