package system

import (
	"fmt"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Ps(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.ProcMutex, s.Procs, pid, s.Host)

	var output string
	output += " PID  TTY    TIME CMD\r\n"
	for _, p := range s.Procs {
		if p.PPID == s.BashPID || p.PPID == 1 {
			output += fmt.Sprintf("%-5d %-5s %5s %s\r\n", p.PID, "tty1", p.StartTime.Format("15:04"), p.Cmd)
		}
	}

	return output, 0
}
