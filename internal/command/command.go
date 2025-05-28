package command

import (
	"fmt"

	"github.com/scGetStuff/gator/internal/config"
	"github.com/scGetStuff/gator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CmdFuncs map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	cmdFunc, ok := c.CmdFuncs[cmd.Name]
	if !ok {
		return fmt.Errorf("'%s' command does not exist", cmd.Name)
	}

	err := cmdFunc(s, cmd)

	return err
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CmdFuncs[name] = f
}
