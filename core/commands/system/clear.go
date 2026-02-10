package system

import (
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Clear(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)
	return "\x1b[H\x1b[2J\x1b[3J", 0
}
