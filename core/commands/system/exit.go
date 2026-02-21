package system

import (
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Exit(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.ProcMutex, s.Procs, pid, s.Host)
	s.Close()
	return "", 0
}
