package espanarusa

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/isalikov/canary.icu-bot/internal/db"
)

const postURL string = "https://espanarusa.com"

type meta struct {
	Pubdate string
	Img     string
	Title   string
	Message string
}

// MakeItemFactory return factory to getting full Post item
func MakeItemFactory() func(id string) (db.Post, error) {
	result := &db.Post{}

	return func(id string) (db.Post, error) {
		hash := md5.Sum([]byte(id))

		URL := fmt.Sprintf("https://espanarusa.com/%v", id)

		r, err := http.NewRequest("GET", URL, nil)

		if err != nil {
			return *result, err
		}

		client := &http.Client{Timeout: time.Second * 30}
		resp, err := client.Do(r)

		if err != nil {
			return *result, err
		}

		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			return *result, err
		}

		img := ""
		message := ""

		article := doc.Find(".js-mediator-article")

		title := doc.Find(".er-page-left").Find("h1").Text()
		pubdate, _ := article.Find("time").Attr("datetime")
		timestamp, _ := time.Parse(time.RFC3339, pubdate)

		article.Find("p").Each(func(i int, c *goquery.Selection) {
			if i == 0 {
				if src, ok := c.Find("img").Attr("src"); ok {
					img = fmt.Sprintf("https://espanarusa.com/%v", src)
				}
			} else {
				if c, err := c.Html(); err == nil {
					message = message + c
				}
			}
		})

		result.Link = URL
		result.Message = message
		result.Title = title
		result.PreviewImage = img
		result.ID = fmt.Sprintf("%x", hash)
		result.Timestamp = timestamp.Unix()

		time.Sleep(time.Millisecond * 50)

		return *result, nil
	}
}
