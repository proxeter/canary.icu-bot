package iestafeta

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/isalikov/news-bot/internal/db"
)

// MakeItemFactory return factory to getting full Post item
func MakeItemFactory(d *data) func(id string) (db.Post, error) {
	payload := &post{}
	result := &db.Post{}

	return func(id string) (db.Post, error) {
		for _, v := range d.Channel.Items {
			if v.GUID == id {
				*payload = v
				break
			}
		}

		if payload.GUID == "" {
			return *result, fmt.Errorf("Can't find item by uid: %v", id)
		}

		previewImage, err := getPreviewImage(id)

		if err != nil {
			result.PreviewImage = ""
		}

		hash := md5.Sum([]byte(*&payload.GUID))

		result.Link = *&payload.Link
		result.Message = *&payload.Description
		result.Title = *&payload.Title
		result.PreviewImage = previewImage
		result.ID = fmt.Sprintf("%x", hash)

		time.Sleep(time.Millisecond * 50)

		return *result, nil
	}
}
