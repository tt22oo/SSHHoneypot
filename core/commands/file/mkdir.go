package file

import (
	"fmt"
	"honeypot/core/filesystem"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Mkdir(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)

	var result string
	for _, name := range args[1:] {
		_, err := filesystem.Fetch(s.Dirs, fmt.Sprintf("%s/%s", s.Path, name))
		if err == nil {
			result += fmt.Sprintf("mkdir: cannot create directory '%s': File exists\r\n", name)
			continue
		}

		s.Entry.Children[name] = &filesystem.Entry{
			Type:     filesystem.Directory,
			Children: make(map[string]*filesystem.Entry),
			Data:     nil,
		}
	}

	err := filesystem.Save(s.Dirs, s.Host)
	if err != nil {
		return "error", 1
	}

	return result, 0
}
