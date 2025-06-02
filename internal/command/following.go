package command

import (
	"context"
	"fmt"

	"github.com/scGetStuff/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
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
