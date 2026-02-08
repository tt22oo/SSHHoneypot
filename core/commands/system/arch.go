package system

import (
	"fmt"
	"honeypot/core/configs"
	"honeypot/core/session"
)

func Arch(s *session.Session, args []string) (string, int) {
	return fmt.Sprintf("%s\r\n", configs.Cfg.System["arch"]), 0
}
