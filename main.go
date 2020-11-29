package main

import (
	"sync"

	"github.com/isalikov/news-bot/feeds/iestafeta"
	"github.com/isalikov/news-bot/feeds/russkoe105fm"
	"github.com/isalikov/news-bot/internal/bot"
	"github.com/isalikov/news-bot/internal/db"
)

func main() {
	var wg sync.WaitGroup

	posts := &[]db.Post{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		feed, _ := russkoe105fm.GetFeed()
		*posts = append(*posts, feed...)
	}()

	go func() {
		defer wg.Done()

		feed, _ := iestafeta.GetFeed()
		*posts = append(*posts, feed...)
	}()

	wg.Wait()
	bot.PushMessages(posts)
}
