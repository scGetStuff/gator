package command

import (
	"context"
	"encoding/json"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("feeds failed: %w", err)
	}

	for _, feed := range feeds {
		data, _ := json.Marshal(feed)
		fmt.Println(string(data))
	}

	return nil
}
