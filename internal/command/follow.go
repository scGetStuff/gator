package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scGetStuff/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("the follow handler expects one argument, the feed URL")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feed, err := s.Db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	return doFollow(s, user.ID, feed.ID)
}

func doFollow(s *State, userID uuid.UUID, fieldID uuid.UUID) error {

	follow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userID,
		FeedID:    fieldID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	// data, _ := json.Marshal(follow)
	// fmt.Println(string(data))

	fmt.Println(follow.FeedName)
	fmt.Println(follow.UserName)

	return nil
}
