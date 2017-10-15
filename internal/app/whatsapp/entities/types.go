package entities

type MessageStatus string

const (
	MessageStatusSent      MessageStatus = "sent"
	MessageStatusDelivered MessageStatus = "delivered"
	MessageStatusRead      MessageStatus = "read"
	MessageStatusFailed    MessageStatus = "failed"
	MessageStatusDeleted   MessageStatus = "deleted"
)

type ComponentType string

const (
	ComponentTypeHeader  ComponentType = "HEADER"
	ComponentTypeFooter  ComponentType = "FOOTER"
	ComponentTypeBody    ComponentType = "BODY"
	ComponentTypeButtons ComponentType = "BUTTONS"
)

type MessageType string

const (
	MessageTypeText        MessageType = "TEXT"
	MessageTypeImage       MessageType = "IMAGE"
	MessageTypeDocument    MessageType = "DOCUMENT"
	MessageTypeReaction    MessageType = "REACTION"
	MessageTypeInteractive MessageType = "INTERACTIVE"
	MessageTypeLocation    MessageType = "LOCATION"
	MessageTypeDate        MessageType = "DATE_TIME"
	MessageTypeCurrency    MessageType = "CURRENCY"
	MessageTypeVideo       MessageType = "VIDEO"
)

type CategoryType string

const (
	CategoryTypeAuth      CategoryType = "AUTHENTICATION"
	CategoryTypeMarketing CategoryType = "MARKETING"
	CategoryTypeUtility   CategoryType = "UTILITY"
)

type ButtonType string

const (
	ButtonTypeURL         ButtonType = "URL"
	ButtonTypePhoneNumber ButtonType = "PHONE_NUMBER"
	ButtonTypeReply       ButtonType = "QUICK_REPLY"
	ButtonTypeCopyCode    ButtonType = "COPY_CODE"
	ButtonTypeCatalog     ButtonType = "CATALOG"
	ButtonTypeOTP         ButtonType = "OTP"
)
