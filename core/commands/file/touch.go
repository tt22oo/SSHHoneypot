package file

import (
	"honeypot/core/filesystem"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Touch(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.ProcMutex, s.Procs, pid, s.Host)
	if len(args) == 2 {
		err := filesystem.Make(s.Dirs, s.Entry, filesystem.File, args[1], "", s.Host)
		if err != nil {
			return "", 1
		}

		return "", 0
	}

	return "touch: missing file operand\r\nTry 'touch --help' for more information.\r\n", 1
}
