package server

import (
	"log"
	"honeypot/core/auth"
	"honeypot/core/commands"
	"honeypot/core/configs"
	"honeypot/core/session/handler"

	"github.com/gliderlabs/ssh"
)

func StartSSH() {
	err := configs.Read()
	if err != nil {
		log.Printf(" \033[31m[ERROR]\033[0m Read Config Error: %s\r\n", err.Error())
		return
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

	log.Printf(" \033[34m[INFO]\033[0m SSH Server Listening on Port %s\r\n", configs.Cfg.Config.Listen)
	err = server.ListenAndServe()
	if err != nil {
		log.Printf(" \033[31m[ERROR]\033[0m SSH Server Error: %s\r\n", err.Error())
	}
}
