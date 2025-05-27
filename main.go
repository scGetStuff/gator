package main

import (
	"fmt"

	"github.com/scGetStuff/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print(err)
	}
	cfg.SetUser("scott")

	cfg, err = config.Read()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(cfg)
}
