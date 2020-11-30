package russkoe105fm

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/isalikov/canary.icu-bot/internal/db"
)

const baseURL = "https://feeds.tildacdn.com/api/getfeed"
const baseOrigin = "http://russkoe-105fm.ru"

// GetFeed return news feed orr error
func GetFeed() ([]db.Post, error) {
	r, err := http.NewRequest("GET", baseURL, nil)

	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	q.Add("feeduid", "988619803892")
	q.Add("size", "20")
	q.Add("slice", "1")
	q.Add("sort[date]", "desc")

	r.URL.RawQuery = q.Encode()

	r.Header.Add("origin", baseOrigin)

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	d := &feedResponse{}
	f := MakeItemFactory(d)

	json.NewDecoder(resp.Body).Decode(d)

	result := make([]db.Post, 0, len(d.Posts))

	for _, p := range d.Posts {
		payload := &db.Post{}

		if isNew, err := payload.Make(p.UID, f, []byte(baseURL)); isNew && err == nil {
			result = append(result, *payload)
		}
	}

	return result, nil
}
