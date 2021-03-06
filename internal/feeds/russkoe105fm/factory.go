package russkoe105fm

import (
	"crypto/md5"
	"fmt"
	"log"
	"time"

	"github.com/isalikov/canary.icu-bot/internal/db"
)

// MakeItemFactory return factory to getting full Post item
func MakeItemFactory(d *feedResponse) func(id string) (db.Post, error) {
	payload := &feedPost{}
	result := &db.Post{}

	return func(id string) (db.Post, error) {
		for _, v := range d.Posts {
			if v.UID == id {
				*payload = v
				break
			}
		}

		if payload.UID == "" {
			return *result, fmt.Errorf("Can't find item by uid: %v", id)
		}

		meta, err := getPostMeta(id)

		if err != nil {
			log.Fatal(err)
		}

		timestamp, err := time.Parse("2006-01-02 15:04:05", meta.Published)

		if err != nil {
			log.Fatal(err)
		}

		hash := md5.Sum([]byte(*&payload.UID))

		result.Link = *&payload.URL
		result.Message = meta.Text
		result.PreviewImage = meta.Image
		result.Title = *&payload.Title
		result.ID = fmt.Sprintf("%x", hash)
		result.Timestamp = timestamp.Unix()

		time.Sleep(time.Millisecond * 50)

		return *result, nil
	}
}
