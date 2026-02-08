package handler

import (
	"honeypot/core/session"
	"honeypot/core/session/shell"
	"log"

	"github.com/gliderlabs/ssh"
)

// main handler
func Handler(s ssh.Session) {
	defer s.Close()

	session, err := session.InitSession(s)
	if err != nil {
		log.Println(err)
		return
	}

	err = shell.Handler(session)
	if err != nil {
		log.Println(err)
	}
}
