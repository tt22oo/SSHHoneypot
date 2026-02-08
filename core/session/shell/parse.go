package shell

import (
	"honeypot/core/commands"
	"honeypot/core/session"
	"honeypot/core/session/stream"
)

func parseInput(input string) []string {
	var (
		quote  rune
		result []string
	)
	buf := ""
	for i, w := range input {
		switch w {
		case ' ':
			if quote == '"' || quote == '\'' {
				buf += string(w)
				continue
			} else if buf == "" {
				continue
			}

			result = append(result, buf)
			buf = ""
		case ';', '>':
			if quote == '"' || quote == '\'' {
				buf += string(w)
				continue
			}

			result = append(result, buf)
			buf = ""
			result = append(result, string(w))
		case '&':
			if quote == '"' || quote == '\'' {
				buf += string(w)
				continue
			} else if i+1 == len([]rune(input)) {
				buf += string(w)
				result = append(result, buf)
				continue
			} else if input[i+1] == '&' {
				if buf != "" {
					result = append(result, buf)
				}
				buf = string(w)
				continue
			} else if buf == "&" {
				buf += string(w)
				result = append(result, buf)
				buf = ""
			}

		case '"', '\'':
			if quote == '"' && w == '"' {
				quote = 0
				result = append(result, buf)
				continue
			} else if quote == '\'' && w == '\'' {
				quote = 0
				result = append(result, buf)
				continue
			}

			if quote == 0 {
				quote = w
			} else {
				buf += string(w)
			}
		default:
			if i+1 == len([]rune(input)) {
				buf += string(w)
				result = append(result, buf)
			}

			buf += string(w)
			continue
		}
	}

	return result
}

func parseShell(s *session.Session, input string) error {
	var (
		command []string
	)
	cmds := parseInput(input)
	for i, cmd := range cmds {
		switch cmd {
		case ";":
			if len(command) == 0 {
				continue
			}

			result, _ := commands.Run(s, command)
			command = make([]string, 0)

			err := stream.Output(s, result)
			if err != nil {
				return err
			}
		case "&&":
			if len(command) == 0 {
				continue
			}

			result, stat := commands.Run(s, command)
			command = nil

			err := stream.Output(s, result)
			if err != nil {
				return err
			}

			if stat != 0 {
				return nil
			}
		default:
			command = append(command, cmd)

			if i+1 == len(cmds) {
				result, _ := commands.Run(s, command)
				command = make([]string, 0)

				err := stream.Output(s, result)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
