package commands

import (
	"fmt"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
	"honeypot/core/utils"
)

func Run(s *session.Session, args []string) (string, int) {
	defer utils.RandomSleep()
	cmd := commandTables[args[0]]
	if cmd == nil {
		return fmt.Sprintf("%s: command not found\r\n", args[0]), 1
	}

	p := &proc.Process{
		PPID: 1,
		User: s.Session.User(),
		Cmd:  args[0],
		Args: args[1:],
	}

	err := p.New(s.Procs, s.Host)
	if err != nil {
		return "", 1
	}

	return cmd(s, args, p.PID)
}
