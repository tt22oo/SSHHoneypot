package file

import (
	"fmt"
	"honeypot/core/filesystem"
	"honeypot/core/session"
)

func Mkdir(s *session.Session, args []string) (string, int) {
	var result string
	for _, name := range args[1:] {
		_, err := filesystem.Fetch(s.Dirs, fmt.Sprintf("%s/%s", s.Path, name))
		if err == nil {
			result += fmt.Sprintf("mkdir: cannot create directory '%s': File exists\r\n", name)
			continue
		}

		s.Entry.Children[name] = &filesystem.Entry{
			Type:     filesystem.TypeDirectory,
			Children: make(map[string]*filesystem.Entry),
			Data:     nil,
		}
	}

	return result, 0
}
