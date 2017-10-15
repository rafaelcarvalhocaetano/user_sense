package entities

import (
	"time"
)

type FlowMessage struct {
	ID                  string     `json:"id,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
	FlowName            string     `json:"flow_name"`
	FlowOrder           int        `json:"flow_order"`
	MessageError        string     `json:"message_error"`
	CurrentTemplateName string     `json:"current_template_name"`
	DefaultTemplateName string     `json:"default_template_name"`
	InteractiveName     string     `json:"interactive_name"`
}

type UserMessage struct {
	ID              string        `json:"id,omitempty"`
	CreatedAt       *time.Time    `json:"created_at,omitempty"`
	UpdatedAt       *time.Time    `json:"updated_at,omitempty"`
	MessageID       string        `json:"message_id"`
	FowID           string        `json:"flow_id"`
	UserName        *string       `json:"user_name,omitempty"`
	UserPhoneNumber string        `json:"user_phone_number,omitempty"`
	Status          MessageStatus `json:"status"`
}
