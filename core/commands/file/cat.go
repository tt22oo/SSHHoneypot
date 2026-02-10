package file

import (
	"errors"
	"fmt"
	"honeypot/core/filesystem"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
	"honeypot/core/session/stream"
	"strings"
)

// cat abc.txt
// cat /tmp/abc.txt
func readFile(s *session.Session, args []string) (string, error) {
	var data string

	if strings.Contains(args[1], "/") {
		entry, err := filesystem.Fetch(s.Dirs, args[1])
		if err != nil {
			return "", errors.New("No such file or directory")
		}
		if entry.Type == filesystem.TypeDirectory {
			return "", errors.New("Is a directory")
		}

		data = *entry.Data
	} else {
		entry := s.Entry.Children[args[1]]
		if entry == nil {
			return "", errors.New("No such file or directory")
		} else if entry.Type == filesystem.TypeDirectory {
			return "", errors.New("Is a directory")
		}

		data = *entry.Data
	}
	return data, nil
}

func handleStdin(s *session.Session) error {
	for {
		output, err := stream.Input(s)
		if err != nil {
			return err
		}

		if output == "\x03" || output == "\x04" {
			break
		}

		err = stream.Output(s, output+"\r\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func Cat(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.Procs, pid, s.Host)

	if len(args) == 1 {
		err := handleStdin(s)
		if err != nil {
			return "", 1
		}
	} else if len(args) == 2 {
		if strings.HasPrefix(args[1], "/proc") {
			output, err := proc.Fetch(s.Procs, args[1])
			if err != nil {
				output := fmt.Sprintf("cat: %s: No such file or directory\r\n", args[1])
				stream.Output(s, output)

				return "", 1
			}

			return output, 0
		} else if strings.HasPrefix(s.Path, "/proc") {
			output, err := proc.Fetch(s.Procs, fmt.Sprintf("%s/%s", s.Path, args[1]))
			if err != nil {
				output := fmt.Sprintf("cat: %s: No such file or directory\r\n", args[1])
				stream.Output(s, output)

				return "", 1
			}

			return output, 0
		} else {
			data, err := readFile(s, args)
			if err != nil {
				return fmt.Sprintf("cat: %s: %s\r\n", args[1], err), 1
			}

			return data, 0
		}
	}

	return "Try 'cat --help' for more information.\r\n", 0
}
