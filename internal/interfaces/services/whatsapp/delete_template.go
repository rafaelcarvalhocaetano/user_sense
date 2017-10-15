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

type DeleteTemplateHttp struct {
	token   string
	url     string
	phoneID string
	wbaID   string
	logger  ports.Logger
}

func NewDeleteTemplateHttp(l ports.Logger) *DeleteTemplateHttp {
	return &DeleteTemplateHttp{
		logger:  l,
		token:   os.Getenv("META_TOKEN"),
		url:     os.Getenv("META_URL"),
		phoneID: os.Getenv("META_PHONE_ID"),
		wbaID:   os.Getenv("META_WBA_ID"),
	}
}

func (meta *DeleteTemplateHttp) DeleteTemplate(name string) error {
	act := "delete_template"

	path := fmt.Sprintf("%v/%v/%v", meta.url, meta.wbaID, "message_templates")
	params := url.Values{}
	params.Add("name", name)

	urlData := fmt.Sprintf("%s?%s", path, params.Encode())
	req, err := http.NewRequest("DELETE", urlData, nil)
	if err != nil {
		meta.logger.Error(act, urlData, http.StatusBadRequest)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", meta.token))

	resp, errDefaultClient := http.DefaultClient.Do(req)
	if errDefaultClient != nil {
		meta.logger.Error(act, urlData, http.StatusBadRequest)
		return errDefaultClient
	}

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusBadRequest {
		responseError := map[string]interface{}{
			"response": resp.Body,
			"status":   resp.Status,
		}
		return errors.New(fmt.Sprintf("%v", responseError))
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		meta.logger.Error(act, resp, http.StatusBadRequest)
		return errors.New("failure request to delete template: code" + strconv.Itoa(resp.StatusCode))
	}

	var response map[string]any
	if errDecoder := json.NewDecoder(resp.Body).Decode(&response); errDecoder != nil {
		meta.logger.Error(act, urlData, http.StatusBadRequest)
		return errDecoder
	}

	meta.logger.Debug(act, urlData, http.StatusBadRequest)
	return nil
}
