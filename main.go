package main

import (
	"sync"

	"github.com/isalikov/canary.icu-bot/internal/bot"
	"github.com/isalikov/canary.icu-bot/internal/db"
	"github.com/isalikov/canary.icu-bot/internal/feeds/espanarusa"
	"github.com/isalikov/canary.icu-bot/internal/feeds/iestafeta"
	"github.com/isalikov/canary.icu-bot/internal/feeds/russkoe105fm"
)

func main() {
	var wg sync.WaitGroup

	posts := make([]db.Post, 0)

	wg.Add(3)

	go func() {
		defer wg.Done()

		feed, _ := russkoe105fm.GetFeed()
		posts = append(posts, feed...)
	}()

	go func() {
		defer wg.Done()

		feed, _ := iestafeta.GetFeed()
		posts = append(posts, feed...)
	}()

	go func() {
		defer wg.Done()

		feed, _ := espanarusa.GetFeed()
		posts = append(posts, feed...)
	}()

	wg.Wait()
	bot.PushMessages(posts)
}
