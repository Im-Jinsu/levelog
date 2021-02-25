package discord

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

// MsgBlock is discord message block
type MsgBlock struct {
	Content  string `json:"content"`
	Username string `json:"username"`
}

// Send : Send Discord
func (data *MsgBlock) Send() error {
	var discordWebhookURL string
	if discordWebhookURL = os.Getenv("LOGGER_ESAY_LEVEL_DISCORDURL"); discordWebhookURL == "" {
		return errors.New("Please check your environment variables. [LOGGER_ESAY_LEVEL_DISCORDURL]")
	}

	pBytes, _ := json.Marshal(data)
	payload := bytes.NewBuffer(pBytes)

	req, err := http.NewRequest("POST", discordWebhookURL, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
