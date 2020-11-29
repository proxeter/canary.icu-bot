package bot

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// TgConfig bot configuration
type TgConfig struct {
	chat string
	key  string
}

// GetConfig return Tg config from os ENV params
func GetConfig() TgConfig {
	config := TgConfig{}

	key := os.Getenv("API_KEY")

	if key == "" {
		panic("Can't get API_KEY")
	}

	chat := os.Getenv("CHANNEL_ID")

	if chat == "" {
		panic("Can't get CHANNEL_ID")
	}

	config.key = key
	config.chat = chat

	return config
}

// PushMessage push message to Tg channel
func PushMessage(message string) error {
	config := GetConfig()

	URL := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", config.key)
	r, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		return err
	}

	q := r.URL.Query()
	q.Add("chat_id", config.chat)
	q.Add("text", message)

	r.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(r)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
