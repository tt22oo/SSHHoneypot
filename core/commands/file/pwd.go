package file

import "honeypot/core/session"

func Pwd(s *session.Session, args []string, pid int) (string, int) {
	return s.Path + "\r\n", 0
}
