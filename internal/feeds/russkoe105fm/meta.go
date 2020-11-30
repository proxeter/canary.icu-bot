package russkoe105fm

import (
	"encoding/json"
	"net/http"
	"time"
)

func getPostMeta(uid string) (postItem, error) {
	result := &postResponse{}

	r, err := http.NewRequest("GET", "https://feeds.tildacdn.com/api/getpost", nil)

	if err != nil {
		return *&result.Post, err
	}

	q := r.URL.Query()
	q.Add("postuid", uid)

	r.URL.RawQuery = q.Encode()

	r.Header.Add("origin", baseOrigin)

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(r)

	if err != nil {
		return *&result.Post, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(result)

	return *&result.Post, nil
}
