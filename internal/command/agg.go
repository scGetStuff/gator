package command

import (
	"context"
	"database/sql"
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
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
		log.Fatal(fmt.Errorf("GetNextFeedToFetch() failed: %w", err))
	}

	err = s.Db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
	})
	if err != nil {
		log.Fatal(fmt.Errorf("MarkFeedFetched() failed: %w", err))
	}

	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Fatal(err)
	}
	_ = printFeed

	err = createPosts(s, &feed, rssFeed)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func createPosts(s *State, feed *database.Feed, rssFeed *rss.RSSFeed) error {

	fmt.Println("creating post records for ", feed.Name)

	for _, item := range rssFeed.Channel.Item {

		// TODO: suposed to add code to handle different formats
		// this one matches all the test data
		pubDate, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			return fmt.Errorf("date format is Bad, M'kay: %w", err)
		}
		// fmt.Println(pubDate.Format(time.RFC1123))

		title := sql.NullString{
			String: html.UnescapeString(item.Title),
			Valid:  true,
		}

		description := sql.NullString{
			String: html.UnescapeString(item.Description),
			Valid:  true,
		}

		post, err := s.Db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			PublishedAt: pubDate,
			Url:         item.Link,
			Title:       title,
			Description: description,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			return fmt.Errorf("CreatePost() failed: %w", err)
		}
		_ = post
		// data, _ := json.Marshal(post)
		// fmt.Println(string(data))
	}

	return nil
}

func printFeed(rssFeed *rss.RSSFeed) {
	fmt.Println()
	fmt.Println()
	fmt.Println(html.UnescapeString(rssFeed.Channel.Title))
	fmt.Println(rssFeed.Channel.Link)
	fmt.Println(html.UnescapeString(rssFeed.Channel.Description))

	for _, item := range rssFeed.Channel.Item {
		fmt.Println()
		fmt.Println(html.UnescapeString(item.Title))
		fmt.Println(item.PubDate)
		fmt.Println(item.Link)
		fmt.Println(html.UnescapeString(item.Description))
	}
}
