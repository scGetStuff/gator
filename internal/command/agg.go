package command

import (
	"context"
	"fmt"
	"html"
	"log"

	"github.com/scGetStuff/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	rssFeed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Fatal(err)
	}

	printStuff(rssFeed)

	return nil
}

func printStuff(rssFeed *rss.RSSFeed) {
	fmt.Println(html.UnescapeString(rssFeed.Channel.Title))
	fmt.Println(rssFeed.Channel.Link)
	fmt.Println(html.UnescapeString(rssFeed.Channel.Description))

	for _, item := range rssFeed.Channel.Item {
		fmt.Println()
		fmt.Println(html.UnescapeString(item.Title))
		fmt.Println(html.UnescapeString(item.PubDate))
		fmt.Println(item.Link)
		fmt.Println(html.UnescapeString(item.Description))
	}
}
