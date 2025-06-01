package command

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	following, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get following: %w", err)
	}

	// data, _ := json.Marshal(follow)
	// fmt.Println(string(data))

	for _, name := range following {
		fmt.Println(name)
	}

	return nil
}
