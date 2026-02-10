package system

import (
	"fmt"
	"honeypot/core/configs"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Arch(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)
	return fmt.Sprintf("%s\r\n", configs.Cfg.System.Arch), 0
}
