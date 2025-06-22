package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {

	feed, err := fetchFeed(context.Background(), rssUrl)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}
