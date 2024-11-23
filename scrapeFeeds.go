package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Fetch next Feed")
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("Mark Feed")
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("cant fetch Feed")
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}
	fmt.Println("==============================")

	return nil
}
