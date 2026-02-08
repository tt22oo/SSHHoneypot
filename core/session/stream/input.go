package stream

import (
	"honeypot/core/session"
)

func Input(s *session.Session) (string, error) {
	var result []byte
	buf := make([]byte, 1)

	for {
		_, err := s.Session.Read(buf)
		if err != nil {
			return "", err
		}

		if buf[0] == '\n' || buf[0] == '\r' {
			err = Output(s, "\r\n")
			if err != nil {
				return "", err
			}

			break
		}

		if buf[0] == 0x03 || buf[0] == 0x04 {
			return string(buf[0]), nil
		}

		if buf[0] == 127 {
			if len(result) < 1 {
				continue
			}
			result = result[:len(result)-1]

			err = Output(s, "\b \b")
			if err != nil {
				return "", err
			}

			continue
		}

		_, err = s.Session.Write(buf)
		if err != nil {
			return "", err
		}

		result = append(result, buf[0])
	}

	return string(result), nil
}
