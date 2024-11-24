package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Ephim135/RSS-feed-Aggregator/internal/database"
)

func browse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) == 1 {
		specifiedlimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Conversion Failed %w", err)
		}
		limit = specifiedlimit
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: <optional-Limit-Posts>")
	}
	posts, err := s.db.GetPostsFromUser(context.Background(), database.GetPostsFromUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("cannot get posts from user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
