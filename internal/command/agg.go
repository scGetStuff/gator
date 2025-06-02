package command

import (
	"context"
	"database/sql"
	"fmt"
	"html"
	"log"
	"time"

	"github.com/scGetStuff/gator/internal/database"
	"github.com/scGetStuff/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("the agg handler expects a single argument, time Between Requests")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *State) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("GetNextFeedToFetch() failed: %w", err)
	}

	err = s.Db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
	})
	if err != nil {
		return fmt.Errorf("MarkFeedFetched() failed: %w", err)
	}

	// fmt.Println(feed.Name)
	// rssFeed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Fatal(err)
	}
	printStuff(rssFeed)

	return nil
}

func printStuff(rssFeed *rss.RSSFeed) {
	fmt.Println()
	fmt.Println()
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
