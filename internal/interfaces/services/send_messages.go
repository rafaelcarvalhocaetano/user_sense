package services

import (
	"botwhatsapp/internal/infra/ports"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/k0kubun/pp/v3"
	"io"
	"net/http"
	"os"
)

type SendMessage struct {
	token   string
	url     string
	phoneID string
	wbaID   string
	logger  ports.Logger
}

func NewSendMessage(l ports.Logger) *SendMessage {
	return &SendMessage{
		logger:  l,
		token:   os.Getenv("META_TOKEN"),
		url:     os.Getenv("META_URL"),
		phoneID: os.Getenv("META_PHONE_ID"),
		wbaID:   os.Getenv("META_WBA_ID"),
	}
}

func (meta *SendMessage) Send(p string, data any) (*string, error) {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		meta.logger.Error("dispatcher", data, http.StatusBadRequest)
		return nil, err
	}

	payload := map[string]interface{}{
		"data": jsonValue,
		"path": p,
	}

	path := fmt.Sprintf("%s/%v/%v", meta.url, meta.phoneID, p)
	req, err := http.NewRequest("POST", path, bytes.NewReader(jsonValue))
	if err != nil {
		meta.logger.Error("dispatcher", payload, http.StatusBadRequest)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", meta.token))
	req.ContentLength = int64(len(jsonValue))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		meta.logger.Error("dispatcher", err.Error(), http.StatusBadRequest)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	//if resp.StatusCode != http.StatusOK {
	//	meta.logger.Error("dispatcher", payload, http.StatusBadRequest)
	//	return nil, errors.New("invalid status code " + strconv.Itoa(resp.StatusCode))
	//}

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusBadRequest {
		return nil, errors.New(fmt.Sprintf("services/whatsapp/send_messages/status: %v", resp.Status))
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		meta.logger.Error("dispatcher", payload, http.StatusBadRequest)
		return nil, err
	}

	payload["data"] = response
	meta.logger.Debug("dispatcher", payload, http.StatusBadRequest)

	_, _ = pp.Println(response)
	id, _ := extractID(response)

	return &id, nil
}

func extractID(data map[string]interface{}) (string, error) {
	messages, ok := data["messages"].([]interface{})
	if !ok || len(messages) == 0 {
		return "", fmt.Errorf("messages array not found or empty")
	}

	message, ok := messages[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unable to parse message object")
	}

	id, ok := message["id"].(string)
	if !ok {
		return "", fmt.Errorf("message ID not found or not a string")
	}

	return id, nil
}
