package server

import (
	"fmt"
	"honeypot/core/auth"
	"honeypot/core/commands"
	"honeypot/core/configs"
	"honeypot/core/session/handler"
	"log"

	"github.com/gliderlabs/ssh"
)

func StartSSH() error {
	err := configs.Read()
	if err != nil {
		return fmt.Errorf("Read Config Error: %s", err.Error())
	}

	commands.Init()

	server := &ssh.Server{
		Addr:    configs.Cfg.Config.Listen,
		Handler: handler.Handler,
		Version: configs.Cfg.Config.Banner,
	}

	if configs.Cfg.Config.Auth.Auth {
		server.PasswordHandler = auth.Auth
		log.Println(" \033[34m[INFO]\033[0m Password Authentication Enabled")
	}

	log.Printf(" \033[34m[INFO]\033[0m Server Listening on Port %s\r\n", configs.Cfg.Config.Listen)
	return fmt.Errorf("SSH Server Error: %s", server.ListenAndServe().Error())
}
