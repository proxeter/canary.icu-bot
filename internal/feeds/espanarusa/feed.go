package espanarusa

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/isalikov/canary.icu-bot/internal/db"
)

const baseURL = "https://espanarusa.com/ru/news/list"

// GetFeed return news feed orr error
func GetFeed() ([]db.Post, error) {
	r, err := http.NewRequest("GET", baseURL, nil)

	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	q.Add("tag", "канарские острова")

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	f := MakeItemFactory()
	result := make([]db.Post, 0)

	doc.Find(".er-object-brief").Each(func(i int, s *goquery.Selection) {
		id, ok := s.First().Attr("href")
		payload := &db.Post{}

		if isNew, err := payload.Make(id, f, []byte(baseURL)); ok && isNew && err == nil {
			result = append(result, *payload)
		}
	})

	return result, nil
}
