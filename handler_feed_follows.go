package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/linus5304/gator/internal/database"
)

func handleFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <feed url>", cmd.name)
	}
	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not find feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed follow: %w", err)
	}
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handleGetFeedFollowsForUser(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not find feeds for user: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feed follows for this user")
		return nil
	}
	fmt.Printf("Feed follows for user %v:\n", user.Name)
	for _, feed := range feeds {
		printFeedFollow(feed.UserName, feed.FeedName)
	}
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User: %v\n", username)
	fmt.Printf("* Feed: %v\n", feedname)
}
