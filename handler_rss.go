package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func handleAggregate(s *state, cmd command) error {
	if len(cmd.args) < 1 || len(cmd.args) > 2 {
		return fmt.Errorf("usage: %v <time_between_requests>", cmd.name)
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("could not get next feed to fetch %s: %s", feed.Name, err)
		return
	}
	_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("could not mark feed %s as fetched: %s", feed.Name, err)
		return
	}
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("could not fetch feed %s: %s", feed.Name, err)
		return
	}
	printAggregatedFeed(rssFeed)
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

func printAggregatedFeed(feed *RSSFeed) {
	for _, item := range feed.Channel.Item {
		fmt.Printf("%v\n", item.Title)
	}
}
