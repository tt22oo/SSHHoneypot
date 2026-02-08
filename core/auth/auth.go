package auth

import (
	"honeypot/core/auth/log"
	"honeypot/core/configs"
	"honeypot/core/utils"

	"github.com/gliderlabs/ssh"
)

func Auth(ctx ssh.Context, password string) bool {
	defer utils.RandomSleep()

	log.Add(ctx.RemoteAddr().String(), ctx.User(), password)

	cfg := configs.Cfg.Config.Auth
	user := cfg.Username == "*" || cfg.Username == ctx.User()
	pass := cfg.Password == "*" || cfg.Password == password

	return user && pass
}
