package shell

import (
	"honeypot/core/session"
	"honeypot/core/session/logger"
	"honeypot/core/session/stream"
)

func Handler(s *session.Session) error {
	defer s.Close()
	for {
		err := writePrompt(s)
		if err != nil {
			return err
		}

		input, err := stream.Input(s)
		if err != nil {
			return err
		}

		logger.Add(logger.Command, input, s.Host, s.ID)

		err = parseShell(s, input)
		if err != nil {
			return err
		}
	}
}
