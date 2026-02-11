package file

import (
	"fmt"
	"honeypot/core/filesystem"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
	"strings"
)

func Cd(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)
	if len(args) == 1 {
		s.Entry, _ = filesystem.Fetch(s.Dirs, "/root")
		s.Path = "/root"

		return "", 0
	} else if len(args) > 2 {
		return "-bash: cd: too many arguments\r\n", 1
	}

	if args[1] == "/proc" {
		s.Path = "/proc"
		s.Entry = &filesystem.Entry{
			Type: filesystem.Directory,
		}

		return "", 0
	}

	var path string
	if strings.HasPrefix(args[1], "/") {
		path = args[1]
	} else {
		path = fmt.Sprintf("%s/%s", s.Path, args[1])
	}

	entry, err := filesystem.Fetch(s.Dirs, path)
	if err != nil {
		return fmt.Sprintf("-bash: cd: %s: %s\r\n", args[1], err), 1
	} else if entry.Type != filesystem.Directory {
		return fmt.Sprintf("-bash: cd: %s: Not a directory\r\n", args[1]), 1
	}

	s.Path = path
	s.Entry = entry

	return "", 0
}
