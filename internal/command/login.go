package command

import (
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	err := s.Cfg.SetUser(cmd.Args[0])
	if err == nil {
		fmt.Println("the user has been set")
	}

	return err
}
