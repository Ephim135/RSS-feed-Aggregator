package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Ephim135/RSS-feed-Aggregator/internal/database"
	"github.com/google/uuid"
)

func follow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.name)
	}
	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("cant get feed by url")
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

	fmt.Printf("Feed Name: %v\n", feedfollow.FeedName)
	fmt.Printf("User Name: %v\n", feedfollow.UserName)

	return nil
}

func following(s *state, cmd command, user database.User) error {
	followRows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, follow := range followRows {
		fmt.Printf("Feed Names: %v\n", follow.FeedName)
	}
	fmt.Println("======================")
	return nil
}

func unfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.name)
	}
	url := cmd.args[0]

	err := s.db.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("Failed deleting follow")
	}
	return nil
}
