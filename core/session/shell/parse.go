package shell

import (
	"honeypot/core/commands"
	"honeypot/core/session"
	"honeypot/core/session/stream"
)

func parseInput(input string) []string {
	var result []string
	quote, buf := false, ""
	for i, w := range input {
		switch w {
		case ' ':
			if quote {
				buf += string(w)
				continue
			} else if buf == "" {
				continue
			}

			result = append(result, buf)
			buf = ""
		case ';', '>':
			if quote {
				buf += string(w)
				continue
			}

			result = append(result, buf)
			buf = ""
			result = append(result, string(w))
		case '&':
			if quote {
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
			if quote {
				quote = false
				result = append(result, buf)

				continue
			}

			quote = true
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
		command  []string
		lastStat int
	)
	cmds := parseInput(input)
	for i, cmd := range cmds {
		switch cmd {
		case ";":
			if len(command) == 0 {
				continue
			}

			result, stat := commands.Run(s, command)
			command = make([]string, 0)
			lastStat = stat

			err := stream.Output(s, result)
			if err != nil {
				return err
			}
		case "&&":
			if len(command) == 0 {
				continue
			} else if lastStat != 0 {
				break
			}

			result, stat := commands.Run(s, command)
			command = make([]string, 0)
			lastStat = stat

			err := stream.Output(s, result)
			if err != nil {
				return err
			}
		default:
			command = append(command, cmd)

			if i+1 == len(cmds) {
				result, stat := commands.Run(s, command)
				command = make([]string, 0)
				lastStat = stat

				err := stream.Output(s, result)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
