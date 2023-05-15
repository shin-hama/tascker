package redmine

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func Feed() {
	URL := os.Getenv("FEED_URL")
	fmt.Println(URL)

	fp := gofeed.NewParser()
	// 社内サーバーなので Proxy を削除する
	fp.Client = ClientWithoutProxy()

	feed, err := fp.ParseURL(URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(feed.Title)
	fmt.Println(feed.FeedType, feed.FeedVersion)
	for _, item := range feed.Items {
		if item == nil {
			break
		}
		fmt.Println(item.Title)
		fmt.Println("\t->", item.Link)
		fmt.Println("\t->", item.PublishedParsed.Format(time.RFC3339))
	}
}

func ClientWithoutProxy() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: nil, // Proxy を削除する場合は nil をセット
		},
	}

	return client
}
