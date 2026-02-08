package system

import "honeypot/core/session"

func Clear(s *session.Session, args []string) (string, int) {
	return "\x1b[H\x1b[2J\x1b[3J", 0
}
