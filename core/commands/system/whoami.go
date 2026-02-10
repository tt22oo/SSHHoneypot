package system

import (
	"fmt"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Whoami(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)
	return fmt.Sprintf("%s\r\n", s.Session.User()), 0
}
