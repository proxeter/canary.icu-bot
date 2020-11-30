package bot

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/isalikov/canary.icu-bot/internal/db"
)

// TgConfig bot configuration
type TgConfig struct {
	chat     string
	key      string
	debounce int
}

func getConfig() TgConfig {
	config := TgConfig{}

	key := os.Getenv("API_KEY")

	if key == "" {
		panic("Can't get API_KEY")
	}

	chat := os.Getenv("CHANNEL_ID")

	if chat == "" {
		panic("Can't get CHANNEL_ID")
	}

	debounce := os.Getenv("DEBOUNCE")

	if value, err := strconv.Atoi(debounce); err != nil {
		config.debounce = value
	} else {
		config.debounce = 2000
	}

	config.key = key
	config.chat = chat

	return config
}

func pushMessage(message string) error {
	fmt.Printf("Sending: %v\n", message)

	config := getConfig()

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

// PushMessages push messages to Tg channel
func PushMessages(posts []db.Post) {
	config := getConfig()

	sort.SliceStable(posts, func(a, b int) bool {
		return posts[a].Timestamp < posts[b].Timestamp
	})

	for _, post := range posts {
		message := fmt.Sprintf("https://canary.icu/v/%v", post.ID)

		pushMessage(message)
		time.Sleep(time.Duration(config.debounce) * time.Millisecond)
	}
}
