package system

import (
	"fmt"
	"honeypot/core/configs"
	"honeypot/core/session"
)

func Uname(s *session.Session, args []string) (string, int) {
	if len(args) < 2 {
		return "Linux\r\n", 0
	}

	switch args[1] {
	case "-m":
		return fmt.Sprintf("%s\r\n", configs.Cfg.System["arch"]), 0
	default:
		return "Try 'uname --help' for more information.\r\n", 0
	}
}
