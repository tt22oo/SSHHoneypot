package system

import (
	"fmt"
	"honeypot/core/session"
)

func Whoami(s *session.Session, args []string) (string, int) {
	return fmt.Sprintf("%s\r\n", s.Session.User()), 0
}
