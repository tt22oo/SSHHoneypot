package commands

import (
	"fmt"
	"honeypot/core/session"
	"honeypot/core/utils"
)

func Run(s *session.Session, args []string) (string, int) {
	defer utils.RandomSleep()
	cmd := commandTables[args[0]]
	if cmd == nil {
		return fmt.Sprintf("%s: command not found\r\n", args[0]), 1
	}

	return cmd(s, args)
}
