package entities

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type TemplateModel struct {
	Name      string          `json:"name"`
	Language  string          `json:"language"`
	Category  string          `json:"category"`
	MetaWbaID string          `json:"meta_wba_id"`
	Component *ComponentModel `json:"component"`
}

type TemplateComponents struct {
	Name       string            `json:"name"`
	Language   string            `json:"language"`
	Category   string            `json:"category"`
	Components []*ComponentModel `json:"components"`
}

type ComponentModel struct {
	Header  *HeaderModel  `json:"header,omitempty"`
	Body    *BodyModel    `json:"body"`
	Footer  *FooterModel  `json:"footer,omitempty"`
	Buttons *ButtonsModel `json:"buttons,omitempty"`
}

type HeaderModel struct {
	Type    *string  `json:"type"`
	Format  *string  `json:"format,omitempty"`
	Text    *string  `json:"text,omitempty"`
	Example *Exemplo `json:"example,omitempty"`
}

type BodyModel struct {
	Type    *string  `json:"type"`
	Text    *string  `json:"text"`
	Example *Exemplo `json:"example,omitempty"`
}

type FooterModel struct {
	Type    *string  `json:"type"`
	Text    string   `json:"text"`
	Example *Exemplo `json:"example,omitempty"`
}

type Exemplo struct {
	HeaderHandle []string   `json:"header_handle,omitempty"`
	HeaderText   []string   `json:"header_text,omitempty"`
	HeaderUrl    []string   `json:"header_url,omitempty"`
	BodyText     [][]string `json:"body_text,omitempty"`
}

type ButtonsModel struct {
	Type    *string            `json:"type"`
	Buttons []*ItemButtonModel `json:"buttons"`
	Example *Exemplo           `json:"example,omitempty"`
}

type ItemButtonModel struct {
	Type        string    `json:"type"`
	Text        string    `json:"text,omitempty"`
	PhoneNumber *string   `json:"phone_number,omitempty"`
	Url         *string   `json:"url,omitempty"`
	OtpType     *string   `json:"otp_type,omitempty"`
	Example     *[]string `json:"example,omitempty"`
}

type MetaPersistence struct {
	Name       string `json:"name"`
	Language   string `json:"language"`
	Category   string `json:"category"`
	Components []any  `json:"components"`
}

func NewTemplateModel(t *TemplateModel) (*TemplateModel, error) {
	return t.validate()
}

func (t *TemplateModel) validate() (*TemplateModel, error) {
	t.Language = "pt_BR"
	var listError []string
	c := strings.ToUpper(t.Category)
	switch CategoryType(c) {
	case CategoryTypeAuth, CategoryTypeMarketing, CategoryTypeUtility:
	default:
		listError = append(listError, "invalid category")
	}
	if t.Name == "" {
		listError = append(listError, "name is required")
	}
	if t.Component == nil {
		listError = append(listError, "component is required")
	}
	if len(listError) > 0 {
		return nil, errors.New(strings.Join(listError, ","))
	}
	t.Category = c

	return t, nil
}

func NewHeaderModel(i *HeaderModel) (*HeaderModel, error) {
	return i.validate()
}

func (h *HeaderModel) validate() (*HeaderModel, error) {
	var listError []string
	*h.Format = strings.ToUpper(*h.Format)
	*h.Type = strings.ToUpper(*h.Type)
	if *h.Type != "HEADER" {
		listError = append(listError, fmt.Sprintf("invalid header type (%v)", *h.Type))
	}

	switch MessageType(*h.Format) {
	case MessageTypeText, MessageTypeImage, MessageTypeDocument, MessageTypeReaction, MessageTypeInteractive, MessageTypeLocation:
	default:
		listError = append(listError, fmt.Sprintf("header format (%v) invalid", *h.Type))
	}

	if MessageType(*h.Type) == MessageTypeDocument && h.Example != nil && len(h.Example.HeaderHandle) == 0 {
		listError = append(listError, "invalid header example handle for document type")
	}
	var txt *string
	if h.Text != nil {
		txt, _ = replaceText(*h.Text)
	}

	if h.Example != nil {
		if h.Example.HeaderText == nil && h.Example.HeaderHandle == nil {
			h.Example = nil
		}
	}
	if len(listError) > 0 {
		return nil, errors.New(strings.Join(listError, ","))
	}
	h.Text = txt
	return h, nil
}

func NewBodyModel(i *BodyModel) (*BodyModel, error) {
	return i.validate()
}

func (b *BodyModel) validate() (*BodyModel, error) {
	var listError []string

	*b.Type = strings.ToUpper(*b.Type)
	if *b.Type != "BODY" {
		listError = append(listError, "invalid type body")
	}
	var count int
	b.Text, count = replaceText(*b.Text)

	if b.Example != nil {
		for _, exc := range b.Example.BodyText {
			if count != len(exc) {
				listError = append(listError, fmt.Sprintf("numero de parametros diferente %v", count))
			}
		}
	}

	if len(listError) > 0 {
		return nil, errors.New(strings.Join(listError, ","))
	}

	return b, nil
}

func NewFooterModel(i *FooterModel) (*FooterModel, error) {
	return i.validate()
}

func (f *FooterModel) validate() (*FooterModel, error) {
	*f.Type = strings.ToUpper(*f.Type)
	if *f.Type != "FOOTER" {
		return nil, errors.New(fmt.Sprintf("invalid type%s", *f.Type))
	}
	return f, nil
}

// Button text can't have more than 25 characters."
func NewButtonModel(i *ButtonsModel) (*ButtonsModel, error) {
	return i.validate()
}

func (b *ButtonsModel) validate() (*ButtonsModel, error) {
	var item []*ItemButtonModel
	var listError []string
	t := strings.ToUpper(*b.Type)
	if t != "BUTTONS" {
		listError = append(listError, fmt.Sprintf("invalid type%s", t))
	}

	for _, btn := range b.Buttons {
		btns, err := NewItemButtonModel(btn)
		if err != nil {
			listError = append(listError, fmt.Sprintf("button item error: %v", err.Error()))
		}
		item = append(item, btns)
	}

	if len(listError) > 0 {
		return nil, errors.New(strings.Join(listError, ","))
	}

	b.Type = &t
	b.Buttons = item
	return b, nil
}

func NewItemButtonModel(i *ItemButtonModel) (*ItemButtonModel, error) {
	return i.validate()
}

func (it *ItemButtonModel) validate() (*ItemButtonModel, error) {
	var listError []string
	tup := strings.ToUpper(it.Type)
	//var countParams int
	switch ButtonType(tup) {
	case ButtonTypeURL:
		it.PhoneNumber = nil
		it.OtpType = nil
		it.Url, _ = replaceText(*it.Url)
	case ButtonTypePhoneNumber:
		it.OtpType = nil
		it.Url = nil
		it.PhoneNumber, _ = replaceText(*it.PhoneNumber)
	case ButtonTypeReply, ButtonTypeCatalog:
		it.OtpType = nil
		it.Url = nil
		it.PhoneNumber = nil
	case ButtonTypeOTP:
		it.Url = nil
		it.PhoneNumber = nil
		it.OtpType, _ = replaceText(*it.OtpType)
	default:
		listError = append(listError, fmt.Sprintf("invalid button type: %v", tup))
	}
	//if it.Example != nil && len(*it.Example) == 0 {
	//	it.Example = replaceText(strings.Join(*it.Example, ","))
	//}

	var newVal *string
	if it.Example != nil {
		for _, exc := range *it.Example {
			newVal, _ = replaceText(exc)
		}
	}

	if newVal != nil {
		it.Example = nil
		x := strings.Split(*newVal, ",")
		it.Example = &x
	}

	if len(listError) > 0 {
		return nil, errors.New(strings.Join(listError, ","))
	}

	it.Type = tup
	return it, nil
}

func (t *TemplateModel) RequestData() *MetaPersistence {
	meta := MetaPersistence{
		Name:       t.Name,
		Language:   t.Language,
		Category:   t.Category,
		Components: make([]any, 0),
	}
	if t.Component.Header != nil {
		meta.Components = append(meta.Components, t.Component.Header)
	}
	if t.Component.Body != nil {
		meta.Components = append(meta.Components, t.Component.Body)
	}
	if t.Component.Footer != nil {
		meta.Components = append(meta.Components, t.Component.Footer)
	}
	if t.Component.Buttons != nil {
		meta.Components = append(meta.Components, t.Component.Buttons)
	}

	return &meta
}

func replaceText(text string) (*string, int) {
	re := regexp.MustCompile(`\[(\d+)\]`)
	var size []int
	ifunc := func(match string) string {
		num := strings.Trim(match, "[]")
		if num != "" {
			size = append(size, len(num))
		}
		return fmt.Sprintf("{{%s}}", num)
	}
	result := re.ReplaceAllStringFunc(text, ifunc)
	return &result, len(size)
}
