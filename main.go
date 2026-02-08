package main

import (
	"honeypot/core/auth"
	"honeypot/core/commands"
	"honeypot/core/configs"
	"honeypot/core/session/handler"

	"github.com/gliderlabs/ssh"
)

func main() {
	err := configs.Read()
	if err != nil {
		panic(err)
	}

	commands.Init()

	server := &ssh.Server{
		Addr:    configs.Cfg.Config.Listen,
		Handler: handler.Handler,
		Version: configs.Cfg.Config.Banner,
	}

	if configs.Cfg.Config.Auth.Auth {
		server.PasswordHandler = auth.Auth
	}

	panic(server.ListenAndServe())
}
