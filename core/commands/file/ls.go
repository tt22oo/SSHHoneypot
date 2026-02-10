package file

import (
	"fmt"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
	"strings"
)

func Ls(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)
	if s.Path == "/proc" {
		output, _ := proc.Fetch(s.Procs, "/proc")
		return output, 0
	}

	var output string
	for name := range s.Entry.Children {
		if len(args) == 2 && args[1] == "-a" {
			output += fmt.Sprintf("%s ", name)
			continue
		} else if strings.HasPrefix(name, ".") {
			continue
		}
		output += fmt.Sprintf("%s ", name)
	}

	if output != "" {
		output += "\r\n"
	}

	return output, 0
}
