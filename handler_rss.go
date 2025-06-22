package main

import (
	"context"
	"fmt"
)

func handlerAggRSS(s *state, cmd command) error {

	rssFeed, err := fetchFeed(context.Background(), rssUrl)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}

	fmt.Println(rssFeed)

	return nil
}
