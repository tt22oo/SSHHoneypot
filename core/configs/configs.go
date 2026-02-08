package configs

import (
	"encoding/json"
	"os"
)

type Auth struct {
	Auth     bool   `json:"auth"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Listen   string `json:"listen"`
	Banner   string `json:"banner"`
	MaxDelay int    `json:"max_delay"`
	Auth     Auth   `json:"auth"`
}

type Configs struct {
	Config Config            `json:"configs"`
	System map[string]string `json:"system"`
}

var Cfg Configs

func Read() error {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var config Configs
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return err
	}

	Cfg = config

	return nil
}
