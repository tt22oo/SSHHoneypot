package system

import (
	"fmt"
	"honeypot/core/session"
)

func Id(s *session.Session, args []string) (string, int) {
	user := s.Session.User()
	return fmt.Sprintf("uid=0(%s) gid=0(%s) groups=0(%s)\r\n", user, user, user), 0
}
