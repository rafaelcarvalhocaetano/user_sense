package whatsapp

import (
	"botwhatsapp/internal/infra/ports"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RegisteTemplateHttp struct {
	token   string
	url     string
	phoneID string
	wbaID   string
	logger  ports.Logger
}

func NewRegisteTemplateHttp(l ports.Logger) *RegisteTemplateHttp {
	return &RegisteTemplateHttp{
		logger:  l,
		token:   os.Getenv("META_TOKEN"),
		url:     os.Getenv("META_URL"),
		phoneID: os.Getenv("META_PHONE_ID"),
		wbaID:   os.Getenv("META_WBA_ID"),
	}
}
func (meta *RegisteTemplateHttp) Register(p string, data any) (map[string]any, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		meta.logger.Error("dispatcher", data, http.StatusBadRequest)
		return nil, err
	}

	payload := map[string]interface{}{
		"data": string(jsonData),
		"path": p,
	}

	fmt.Println("payload: ", payload)

	path := fmt.Sprintf("%s/%v/%v", meta.url, meta.wbaID, p)
	req, errR := http.NewRequest("POST", path, bytes.NewReader(jsonData))
	if errR != nil {
		meta.logger.Error("dispatcher", payload, http.StatusBadRequest)
		return nil, errR
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", meta.token))
	req.ContentLength = int64(len(jsonData))

	resp, errDefaultClient := http.DefaultClient.Do(req)
	if errDefaultClient != nil {
		meta.logger.Error("dispatcher", payload, http.StatusBadRequest)
		return nil, errDefaultClient
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusBadRequest {
		responseError := map[string]any{
			"code":   resp.StatusCode,
			"status": resp.Status,
		}
		return nil, errors.New(fmt.Sprintf("%v", responseError))
	}

	var response map[string]interface{}
	if errDecoder := json.NewDecoder(resp.Body).Decode(&response); errDecoder != nil {
		meta.logger.Error("dispatcher", payload, http.StatusBadRequest)
		return nil, errDecoder
	}

	payload["data"] = response
	meta.logger.Debug("dispatcher", payload, http.StatusBadRequest)
	return response, nil
}
