package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ephim135/RSS-feed-Aggregator/internal/database"
	"github.com/google/uuid"
)

func addfeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	feedfollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println("Feed Followed successfully:")
	printFeedFollow(feedfollow.UserName, feedfollow.FeedName)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func feed(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s does not need argument", cmd.name)
	}

	feedrows, err := s.db.Feed(context.Background())
	if err != nil {
		return fmt.Errorf("couldnt read feed data")
	}
	for _, feed := range feedrows {
		fmt.Printf("Feed Name: %v\n", feed.FeedsName)
		fmt.Printf("url: %v\n", feed.Url)
		fmt.Printf("User Name: %v\n", feed.UserName)
		fmt.Printf("===============================\n\n")
	}
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
