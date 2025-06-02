package command

import (
	"context"
	"fmt"

	"github.com/scGetStuff/gator/internal/database"
)

func HandlerUnFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("the unfollow handler expects one argument, the feed URL")
	}

	feed, err := s.Db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	err = s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unfollow failed: %w", err)
	}

	return nil
}
