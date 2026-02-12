package file

import (
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Pwd(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)
	return s.Path + "\r\n", 0
}
