package iestafeta

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/isalikov/news-bot/internal/db"
)

const baseURL = "https://iestafeta.com/feed"

// GetFeed return news feed orr error
func GetFeed() ([]db.Post, error) {
	r, err := http.NewRequest("GET", baseURL, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	d := &data{}
	f := MakeItemFactory(d)

	xml.NewDecoder(resp.Body).Decode(d)

	result := make([]db.Post, 0, len(d.Channel.Items))

	for _, p := range d.Channel.Items {
		payload := &db.Post{}

		if isNew, err := payload.Make(p.UID, f, []byte(baseURL)); isNew && err == nil {
			result = append(result, *payload)
		}
	}

	return result, nil
}
