package entities

import (
	"fmt"
	"testing"
)

func TestValidateItemButtonModel(t *testing.T) {
	btn := &ItemButtonModel{
		Type:        "url",
		Url:         strPtr("https://teste.com[1]-data-[2]"),
		PhoneNumber: strPtr("123456789"),
		OtpType:     strPtr("OTP123"),
	}
	resp, err := NewItemButtonModel(btn)
	if err != nil {
		t.Errorf("error no valor: %v", err.Error())
	}
	t.Log(resp)
}

func TestValidateItemButtonMainModel(t *testing.T) {
	btn := &ButtonsModel{
		Type:    strPtr("BUTTONS"),
		Buttons: make([]*ItemButtonModel, 0),
	}

	btn.Buttons = append(btn.Buttons, &ItemButtonModel{
		Type:        "url",
		Url:         strPtr("https://teste.com[1]-data-[2]"),
		PhoneNumber: strPtr("123456789"),
		OtpType:     strPtr("OTP123"),
	})
	_, err := NewButtonModel(btn)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestValidateTemplate(t *testing.T) {
	tm, err := NewTemplateModel(&TemplateModel{
		Name:     "identidade",
		Category: "utility",
		Component: &ComponentModel{
			Header:  nil,
			Body:    nil,
			Footer:  nil,
			Buttons: nil,
		},
	})

	if err != nil {
		t.Errorf("error no valor: %v", err.Error())
	}
	fmt.Println(tm)
}

func strPtr(s string) *string {
	return &s
}
