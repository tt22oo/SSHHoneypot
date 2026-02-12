package system

import (
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
	"strconv"
	"time"
)

func Sleep(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.ProcMutex, s.Procs, pid, s.Host)
	if len(args) != 2 {
		return "Try 'sleep --help' for more information.\r\n", 1
	}

	t, err := strconv.Atoi(args[1])
	if err != nil {
		return "Try 'sleep --help' for more information.\r\n", 1
	}

	time.Sleep(time.Duration(t) * time.Second)

	return "", 0
}
