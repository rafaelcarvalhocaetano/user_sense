package channel

import (
	"botwhatsapp/internal/interfaces/webhooks/model"
	"sync"
)

type ChannelStatusMessage struct {
}

func NewChannelStatusMessage() *ChannelStatusMessage {
	return &ChannelStatusMessage{}
}

func (chs *ChannelStatusMessage) ChannelStatusMessage(statuses <-chan model.Status) {
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		//for stt := range statuses {
		//	wg.Add(1)
		//	name := fmt.Sprintf("S/N")
		//	mid := os.Getenv("META_WBA_ID")
		//	msg := "S/M"
		//	_ = chs.repository.DojoSaveUserMessage(&name, stt.RecipientID, &msg, &stt.Status, &stt.ID, &mid)
		//}
	}()
	go func() { wg.Wait() }()
}
