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

	var server *ssh.Server
	if configs.Cfg.Config.Auth.Auth {
		server = &ssh.Server{
			Addr:            configs.Cfg.Config.Listen,
			PasswordHandler: auth.Auth,
			Handler:         handler.Handler,
			Version:         configs.Cfg.Config.Banner,
		}
	} else {
		server = &ssh.Server{
			Addr:    configs.Cfg.Config.Listen,
			Handler: handler.Handler,
			Version: configs.Cfg.Config.Banner,
		}
	}

	panic(server.ListenAndServe())
}
