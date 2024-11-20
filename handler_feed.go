package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/linus5304/gator/internal/database"
)

func handleAddFeed(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName
	if currentUser == "" {
		return fmt.Errorf("you must be logged in to add a feed")
	}
	user, err := s.db.GetUserByName(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %v <feed name> <feed url>", cmd.name)
	}
	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}
	fmt.Printf("Feed created: %+v\n", feed)

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
	fmt.Printf("Feed followed: %+v\n", feedFollow)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feedsAndUsers, err := s.db.GetFeedsAndUser(context.Background())
	if err != nil {
		return fmt.Errorf("could not find feeds")
	}
	for _, feed := range feedsAndUsers {
		fmt.Printf("%v\n", feed.Name)
		fmt.Printf("%v\n", feed.Url)
		fmt.Printf("%v\n", feed.UserName)
	}
	return nil
}

// func printFeed(feed database.Feed) {
// 	fmt.Printf("%+v\n", feed.Name)
// 	fmt.Printf("%+v\n", feed.Url)
// }
