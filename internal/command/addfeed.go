package command

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scGetStuff/gator/internal/database"
)

func HandlerAddfeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("the addfeed handler expects two arguments, the feed name and URL")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	// fmt.Print(feed)
	// data, _ := json.Marshal(feed)
	// fmt.Println(string(data))

	err = doFollow(s, user.ID, feed.ID)
	if err != nil {
		return err
	}

	return nil
}
