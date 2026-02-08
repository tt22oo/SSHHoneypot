package session

import (
	"fmt"
	"honeypot/core/configs"
	"honeypot/core/filesystem"
	"honeypot/core/session/log"
	"net"
	"os"

	"github.com/gliderlabs/ssh"
)

type Session struct {
	ID      string
	Session ssh.Session
	Host    string
	Path    string
	Entry   *filesystem.Entry
	Dirs    map[string]*filesystem.Entry
}

func newDirs(host string) (*os.File, error) {
	path := fmt.Sprintf("sessions/%s", host)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}

	path += "/dirs.json"
	_, err = os.Stat(path)
	if err != nil {
		data, err := os.ReadFile("configs/dirs.json")
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(path, data, 0644)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.Open(path)
	return f, err
}

func fetchID(host string) (string, error) {
	path := fmt.Sprintf("sessions/%s/", host)
	entries, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	var n int
	for _, entry := range entries {
		if entry.IsDir() {
			n++
		}
	}

	return fmt.Sprintf("%d", n+1), nil
}

func newSession(s ssh.Session) (*Session, *os.File, error) {
	host, _, err := net.SplitHostPort(s.RemoteAddr().String())
	if err != nil {
		return nil, nil, err
	}

	f, err := newDirs(host)
	if err != nil {
		return nil, nil, err
	}

	id, err := fetchID(host)
	if err != nil {
		return nil, nil, err
	}

	return &Session{
		ID:      id,
		Session: s,
		Host:    host,
		Path:    "/root",
	}, f, nil
}

func InitSession(s ssh.Session) (*Session, error) {
	session, f, err := newSession(s)
	if err != nil {
		return nil, err
	}

	session.Dirs, err = filesystem.Parse(f)
	if err != nil {
		return nil, err
	}
	session.Entry = session.Dirs["root"]

	log.Add(log.Connection, session.Host, session.Host, session.ID)

	banner, err := configs.ReadBanner()
	if err != nil {
		return nil, err
	}

	_, err = s.Write([]byte("\x1b[H\x1b[2J\x1b[3J"))
	if err != nil {
		return nil, err
	}

	_, err = s.Write(banner)
	if err != nil {
		return nil, err
	}

	return session, nil
}
