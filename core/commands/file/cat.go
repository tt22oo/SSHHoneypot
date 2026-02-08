package file

import (
	"errors"
	"fmt"
	"honeypot/core/filesystem"
	"honeypot/core/session"
	"honeypot/core/session/stream"
	"log"
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

func Cat(s *session.Session, args []string) (string, int) {
	switch len(args) {
	case 1:
		err := handleStdin(s)
		if err != nil {
			log.Println(err)
			return "", 1
		}
	case 2:
		data, err := readFile(s, args)
		if err != nil {
			return fmt.Sprintf("cat: %s: %s\r\n", args[1], err), 1
		}

		return data, 0
	default:
		return "", 0
	}

	return "", 0
}
