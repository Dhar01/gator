package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
)

func HandlerFollow(s *commands.State, cmd commands.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("USAGE: %s <URL>\n", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed not found for URL: %s\n", err)
		}
		return err
	}

	followFeed, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w\n", err)
	}

	fmt.Println("Feed follow created!")

	fmt.Printf("Feed Name: %s\n", followFeed.FeedName)
	fmt.Printf("User Name: %s\n", followFeed.UserName)

	return nil
}
