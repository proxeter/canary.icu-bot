package russkoe105fm

import (
	"crypto/md5"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/isalikov/news-bot/internal/db"
)

func removeHTMLTag(in string) string {
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(in, -1)

	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			in = strings.ReplaceAll(in, group, "")
		}
	}
	return in
}

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
		result.Message = removeHTMLTag(meta.Text)
		result.PreviewImage = meta.Image
		result.Title = *&payload.Title
		result.ID = fmt.Sprintf("%x", hash)
		result.Timestamp = timestamp.Unix()

		time.Sleep(time.Millisecond * 50)

		return *result, nil
	}
}
