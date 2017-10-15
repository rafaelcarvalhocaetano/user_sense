package whatsapp

import (
	"botwhatsapp/internal/infra/ports"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type TemplateHttp struct {
	token   string
	url     string
	phoneID string
	wbaID   string
	logger  ports.Logger
}

func NewTemplateService(l ports.Logger) *TemplateHttp {
	return &TemplateHttp{
		logger:  l,
		token:   os.Getenv("META_TOKEN"),
		url:     os.Getenv("META_URL"),
		phoneID: os.Getenv("META_PHONE_ID"),
		wbaID:   os.Getenv("META_WBA_ID"),
	}
}

func (meta *TemplateHttp) GetDataModel(p string, data map[string]string) (map[string]any, error) {
	payload := map[string]any{
		"data": data,
		"path": p,
	}
	path := fmt.Sprintf("%s/%v/%v", meta.url, meta.wbaID, p)
	parsedURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := parsedURL.Query()
	for key, value := range data {
		query.Set(key, value)
	}
	parsedURL.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", meta.token))
	req.ContentLength = int64(len(query))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusBadRequest {
		responseError := map[string]interface{}{
			"response": resp.Body,
			"status":   resp.Status,
		}
		return nil, errors.New(fmt.Sprintf("%v", responseError))
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code " + strconv.Itoa(resp.StatusCode))
	}

	var response map[string]any
	if errDecoder := json.NewDecoder(resp.Body).Decode(&response); errDecoder != nil {
		return nil, errDecoder
	}

	delete(response, "paging")
	payload["data"] = response
	//meta.logger.Debug("dispatcher", payload, http.StatusBadRequest)
	return response, nil
}
