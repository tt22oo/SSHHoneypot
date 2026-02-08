package file

import (
	"fmt"
	"honeypot/core/session"
	"strings"
)

func Ls(s *session.Session, args []string) (string, int) {
	var output string
	for name := range s.Entry.Children {
		if strings.HasPrefix(name, ".") {
			continue
		}
		output += fmt.Sprintf("%s ", name)
	}

	if output != "" {
		output += "\r\n"
	}

	return output, 0
}
