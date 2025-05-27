package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	c := Config{}

	cfgFile, err := getConfigFilePath()
	if err != nil {
		return c, err
	}

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

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := write(*c)

	return err
}

func write(c Config) error {
	stuff, err := json.Marshal(c)
	if err != nil {
		return err
	}

	cfgFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	os.WriteFile(cfgFile, stuff, 0666)

	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + "/.gatorconfig.json", nil
}
