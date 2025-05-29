package command

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("users failed: %w", err)
	}

	for _, user := range users {
		out := user
		if s.Cfg.CurrentUserName == user {
			out += " (current)"
		}

		fmt.Println(out)
	}

	return nil
}
