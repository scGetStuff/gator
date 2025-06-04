package command

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/scGetStuff/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {

	var limit int32 = 2
	if len(cmd.Args) > 0 {
		x, _ := strconv.Atoi(cmd.Args[0])
		if x > 0 {
			limit = int32(x)
		}
	}
	// fmt.Println(limit)

	posts, err := s.Db.GetUserPosts(context.Background(),
		database.GetUserPostsParams{UserID: user.ID, Limit: limit})
	if err != nil {
		return fmt.Errorf("GetUserPosts() failed: %w", err)
	}

	for _, post := range posts {
		printPost(&post)
	}

	return nil
}

func printPost(post *database.Post) {
	// data, _ := json.Marshal(post)
	// fmt.Println(string(data))

	fmt.Println()
	fmt.Println(post.Title.String)
	fmt.Println(post.PublishedAt.Format(time.RFC1123))
	fmt.Println(post.Url)
	fmt.Println(post.Description.String)
}
