package file

import (
	"fmt"
	"honeypot/core/session"
	"strings"
)

func Ls(s *session.Session, args []string) (string, int) {
	var output string
	for name := range s.Entry.Children {
		if len(args) >= 2 {
			switch args[1] {
			case "-a":
				output += fmt.Sprintf("%s ", name)
			default:
				if strings.HasPrefix(name, ".") {
					continue
				}
				output += fmt.Sprintf("%s ", name)
			}
		} else {
			if strings.HasPrefix(name, ".") {
				continue
			}
			output += fmt.Sprintf("%s ", name)
		}
	}

	if output != "" {
		output += "\r\n"
	}

	return output, 0
}
