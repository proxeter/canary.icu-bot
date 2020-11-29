package main

import (
	"fmt"

	"github.com/isalikov/news-bot/feeds/iestafeta"
)

func main() {
	r, _ := iestafeta.GetFeed()

	fmt.Println(r)
}
