package shell

import (
	"honeypot/core/session"
	"honeypot/core/session/log"
	"honeypot/core/session/stream"
)

func Handler(s *session.Session) error {
	for {
		err := writePrompt(s)
		if err != nil {
			return err
		}

		output, err := stream.Input(s)
		if err != nil {
			return err
		}

		log.Add(log.Command, output, s.Host, s.ID)

		err = parseShell(s, output)
		if err != nil {
			return err
		}
	}
}
