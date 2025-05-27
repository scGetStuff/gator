package main

import (
	"fmt"
	"os"

	c "github.com/scGetStuff/gator/internal/command"
	"github.com/scGetStuff/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	// L2
	// cfg.SetUser("scott")
	// cfg, err = config.Read()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cfg)

	// L3
	state := c.State{Cfg: &cfg}
	cmds := initCommandMap()
	if len(os.Args) < 2 {
		fmt.Println("not enough stuff to do any stuff with")
		os.Exit(1)
	}
	command := c.Command{Name: os.Args[1], Args: os.Args[2:]}
	err = cmds.Run(&state, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func initCommandMap() c.Commands {
	cmds := c.Commands{CmdFuncs: map[string]func(*c.State, c.Command) error{}}
	cmds.Register("login", c.HandlerLogin)

	return cmds
}
