package handler

import (
	"honeypot/core/filesystem/proc"
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
		log.Printf(" \033[31m[ERROR]\033[0m Init Session Error: %s\r\n", err.Error())
		return
	}
	defer proc.Delete(session.ProcMutex, session.Procs, session.BashPID, session.Host)

	err = shell.Handler(session)
	if err != nil {
		log.Printf(" \033[31m[ERROR]\033[0m Shell Error: %s\r\n", err.Error())
	}
}
