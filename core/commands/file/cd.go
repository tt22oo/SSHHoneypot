package file

import (
	"fmt"
	"honeypot/core/filesystem"
	"honeypot/core/session"
	"strings"
)

func Cd(s *session.Session, args []string) (string, int) {
	if len(args) <= 1 {
		entry, _ := filesystem.Fetch(s.Dirs, "/root")
		s.Path = "/root"
		s.Entry = entry

		return "", 0
	} else if len(args) > 2 {
		return "-bash: cd: too many arguments\r\n", 1
	}

	var path string
	if strings.Contains(args[1], "/") {
		path = args[1]
	} else {
		path = fmt.Sprintf("%s/%s", s.Path, args[1])
	}

	entry, err := filesystem.Fetch(s.Dirs, path)
	if err != nil {
		return fmt.Sprintf("-bash: cd: %s: %s\r\n", args[1], err), 1
	}

	if entry.Type != filesystem.TypeDirectory {
		return fmt.Sprintf("-bash: cd: %s: Not a directory\r\n", args[1]), 1
	}

	s.Path = path
	s.Entry = entry

	return "", 0
}
