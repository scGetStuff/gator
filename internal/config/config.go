package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var home, _ = os.UserHomeDir()
var cfgFile = home + "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	c := Config{}

	bytes, err := os.ReadFile(cfgFile)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (c *Config) SetUser(name string) {
	c.CurrentUserName = name
	write(*c)
}

func write(c Config) {
	stuff, err := json.Marshal(c)
	if err != nil {
		fmt.Print(err)
		return
	}
	os.WriteFile(cfgFile, stuff, 0666)
}
