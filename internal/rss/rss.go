package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func getBytes(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	data, err := getBytes(feedURL)
	if err != nil {
		return &RSSFeed{}, err
	}

	// fmt.Println(string(data))

	var stuff RSSFeed
	err = xml.Unmarshal(data, &stuff)
	if err != nil {
		return &RSSFeed{}, err
	}

	return &stuff, nil
}
