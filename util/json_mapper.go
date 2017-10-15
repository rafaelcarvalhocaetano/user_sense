package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func ToStruct[T any](data io.ReadCloser) (*T, error) {
	body, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}
	var inputs T
	err = json.Unmarshal(body, &inputs)
	return &inputs, err
}

func DecimalPrecision(valor float64, prefix int) float64 {
	return float64(int(valor*float64(prefix)+0.5)) / float64(prefix)
}

func Converte[T any](data any, output T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.New("json marshal error")
	}

	var person T
	err = json.Unmarshal(jsonData, &person)
	if err != nil {
		return errors.New("json Unmarshal error")
	}
	return nil
}

func ToConvert(value any) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return "", fmt.Errorf("tipo desconhecido: %T", value)
	}
}
