package stream

import "honeypot/core/session"

func Output(s *session.Session, data string) error {
	_, err := s.Session.Write([]byte(data))
	return err
}
