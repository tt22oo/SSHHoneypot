package configs

import (
	"encoding/json"
	"log"
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

type System struct {
	Arch     string `json:"arch"`
	HostName string `json:"host_name"`
}

type Configs struct {
	Config Config `json:"configs"`
	System System `json:"system"`
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

	log.Println(" \033[32m[SUCCESS]\033[0m Configs Read")

	return nil
}
