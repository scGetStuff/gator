package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"

	c "github.com/scGetStuff/gator/internal/command"
	"github.com/scGetStuff/gator/internal/config"
	"github.com/scGetStuff/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	state := &c.State{Db: dbQueries, Cfg: &cfg}

	cmds := initCommandMap()
	if len(os.Args) < 2 {
		log.Fatal("see the README.md for supported commands")
	}
	command := c.Command{Name: strings.ToLower(os.Args[1]), Args: os.Args[2:]}
	err = cmds.Run(state, command)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func initCommandMap() c.Commands {
	cmds := c.Commands{CmdFuncs: map[string]func(*c.State, c.Command) error{}}
	cmds.Register("register", c.HandlerRegister)
	cmds.Register("login", c.HandlerLogin)
	cmds.Register("reset", c.HandlerReset)
	cmds.Register("users", c.HandlerUsers)
	cmds.Register("feeds", c.HandlerFeeds)
	cmds.Register("agg", c.HandlerAgg)

	cmds.Register("addfeed", c.MiddlewareLoggedIn(c.HandlerAddfeed))
	cmds.Register("follow", c.MiddlewareLoggedIn(c.HandlerFollow))
	cmds.Register("unfollow", c.MiddlewareLoggedIn(c.HandlerUnFollow))
	cmds.Register("following", c.MiddlewareLoggedIn(c.HandlerFollowing))
	cmds.Register("browse", c.MiddlewareLoggedIn(c.HandlerBrowse))

	return cmds
}
