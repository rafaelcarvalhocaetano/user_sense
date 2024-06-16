package util

import (
	"encoding/json"
	"io"
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
