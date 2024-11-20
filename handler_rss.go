package main

import (
	"context"
	"fmt"
)

func handleAggregate(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}
	printAggregatedFeed(feed)
	return nil
}

func printAggregatedFeed(feed *RSSFeed) {
	for _, item := range feed.Channel.Item {
		fmt.Printf("%v\n%v\n", item.Title, item.Description)
	}
}
