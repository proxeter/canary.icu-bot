package iestafeta

import (
	"fmt"
	"time"

	"github.com/isalikov/news-bot/internal/db"
)

// MakeItemFactory return factory to getting full Post item
func MakeItemFactory(d *data) func(uid string) (db.Post, error) {
	payload := &post{}
	result := &db.Post{}

	return func(uid string) (db.Post, error) {
		for _, v := range d.Channel.Items {
			if v.UID == uid {
				*payload = v
				break
			}
		}

		if payload.UID == "" {
			return *result, fmt.Errorf("Can't find item by uid: %v", uid)
		}

		previewImage, err := getPreviewImage(uid)

		if err != nil {
			result.PreviewImage = ""
		}

		result.Link = *&payload.Link
		result.Message = *&payload.Message
		result.Title = *&payload.Title
		result.UID = *&payload.UID
		result.PreviewImage = previewImage

		time.Sleep(time.Millisecond * 50)

		return *result, nil
	}
}
