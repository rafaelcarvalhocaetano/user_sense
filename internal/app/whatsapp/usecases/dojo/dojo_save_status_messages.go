package dojo

import (
	"fmt"
	"os"
)

type SaveStatusUserMessage struct {
	db Repository
}

func NewSaveStatusUserMessage(db Repository) *SaveStatusUserMessage {
	return &SaveStatusUserMessage{db: db}
}

func (sts *SaveStatusUserMessage) SaveStatusUserMessage(to string, response map[string]interface{}) error {
	var responseID string
	var status string
	if data, ok := response["data"].(map[string]interface{}); ok {
		if success, ok := data["success"].(map[string]interface{}); ok {
			if messages, ok := success["messages"].([]interface{}); ok {
				if len(messages) > 0 {
					if firstMessage, ok := messages[0].(map[string]interface{}); ok {
						id, _ := firstMessage["id"].(string)
						s, _ := firstMessage["message_status"].(string)
						responseID = id
						status = s
					}
				}
			}
		}
	}

	ld := fmt.Sprintf("user_see: %v", status)
	ln := fmt.Sprintf("MBM")
	recip := os.Getenv("META_WBA_ID")
	_ = sts.db.DojoSaveUserMessage(&ln, &to, &ld, &status, &responseID, &recip)
	return nil
}
