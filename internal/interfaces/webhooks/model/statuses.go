package model

// Statuses mantém você informado sobre o status das mensagens entre você e usuários ou grupos.
type Statuses struct {
	ID           string        `json:"id"`
	RecipientID  string        `json:"recipient_id"`
	Status       string        `json:"status"`
	Timestamp    string        `json:"timestamp"`
	Type         string        `json:"type"`
	Conversation *Conversation `json:"conversation,omitempty"`
	Pricing      *Pricing      `json:"pricing,omitempty"`
}

// Pricing inclui os atributos de faturamento.
type Pricing struct {
	PricingModel string `json:"pricing_model"`
	Billing      bool   `json:"billing"`
	Category     string `json:"category"`
}

// Origin descreve de onde uma conversa se originou.
type Origin struct {
	Type string `json:"type"`
}

// Conversation rastreia os atributos da sua conversa atual.
type Conversation struct {
	ID                  string  `json:"id"`
	Origin              *Origin `json:"origin,omitempty"`
	ExpirationTimestamp *any    `json:"expiration_timestamp,omitempty"`
}
