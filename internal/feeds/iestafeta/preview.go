package iestafeta

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getPreviewImage(uid string) (string, error) {
	r, err := http.Get(uid)

	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	doc, err := goquery.NewDocumentFromReader(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	src, _ := doc.Find(".attachment-main-slider").First().Attr("src")

	return src, nil
}
