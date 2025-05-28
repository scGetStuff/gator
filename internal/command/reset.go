package command

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.Db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("reset failed: %w", err)
	}

	fmt.Println("users table emptied")
	return nil
}
