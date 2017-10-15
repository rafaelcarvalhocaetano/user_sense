package service

import (
	"botwhatsapp/internal/app/whatsapp/entities"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CheckType(value any) string {
	switch v := value.(type) {
	case string:
		if IsDate(v) {
			return string(entities.MessageTypeText)
		} else if IsCurrencyOrNumber(v) {
			return string(entities.MessageTypeCurrency)
		} else if IsCheckURL(v) {
			return string(entities.MessageTypeImage)
		} else {
			return string(entities.MessageTypeText)
		}
	case float64:
		return string(entities.MessageTypeCurrency)
	default:
		return string(entities.MessageTypeText)
	}
}

func IsDate(value string) bool {
	_, err := time.Parse(value, value)
	if err != nil {
		return true
	}

	return false
}

func IsCurrencyOrNumber(value string) bool {
	re := regexp.MustCompile(`^R\$\s?\d{1,tres}(\.\d{tres})*,\d{2}$|^R\$\s?\d+,\d{2}$|^R\$\d+\.\d{2}$`)
	if re.MatchString(value) {
		return true
	}

	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func IsCheckURL(u string) bool {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	return true
}

func MapperExemple(data map[string]any) []any {
	example, ok := data["example"].(map[string]any)
	if !ok {
		return nil
	}

	appendExemplo := func(ex []any, key string) ([]any, bool) {
		value, ok := example[key].([]any)
		if ok {
			if key == "body_text" {
				for _, v := range value {
					ex = append(ex, v)
				}
			} else {
				ex = append(ex, value)
			}
		}

		return ex, ok
	}

	exs := make([]any, 0)
	keys := []string{"header_text", "header_handle", "header_url", "body_text"}

	for _, key := range keys {
		var valid bool
		exs, valid = appendExemplo(exs, key)
		if !valid {
			continue
		}
	}

	return exs
}

func ParseCurrency(value string) (int, error) {
	valorLimpo := strings.Replace(value, "R$", "", 1)
	valorLimpo = strings.Replace(valorLimpo, ".", "", -1)
	valorLimpo = strings.Replace(valorLimpo, ",", ".", 1)
	valorFloat, err := strconv.ParseFloat(valorLimpo, 64)
	if err != nil {
		return 0, err
	}
	return int(valorFloat * 1000), nil
}

func HeaderType(td, value string) bool {
	if td == "image" || td == "document" || td == "video" {
		urlRegex := `^(https?):\/\/[^\s/$.?#].[^\s]*$`
		regex := regexp.MustCompile(urlRegex)

		return regex.MatchString(value)
	}

	return true
}
