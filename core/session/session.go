package session

import (
	"fmt"
	"honeypot/core/configs"
	"honeypot/core/filesystem"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session/logger"
	"net"
	"os"
	"sync"

	"github.com/gliderlabs/ssh"
)

type Session struct {
	ID        string
	BashPID   int
	Session   ssh.Session
	Host      string
	Path      string
	Entry     *filesystem.Entry
	Dirs      map[string]*filesystem.Entry
	Procs     map[int]*proc.Process // key: pid
	ProcMutex *sync.Mutex
}

const (
	sessionPATH string = "sessions/"
	configPATH string = "configs"
	procPATH   string = "/proc"
	procJSON   string = "/procs.json"
	dirJSON    string = "/dirs.json"
)

func (s *Session) copyJSON(src, dst string) (*os.File, error) {
	err := os.MkdirAll(fmt.Sprintf("%s%s", sessionPATH, s.Host), 0755)
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(dst)
	if err != nil {
		data, err := os.ReadFile(src)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(dst, data, 0644)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.Open(dst)
	return f, err
}

func (s *Session) fetchID() (string, error) {
	path := fmt.Sprintf("%s%s/", sessionPATH, s.Host)
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

func (s *Session) Close() {
	defer s.Session.Close()
	proc.Delete(s.ProcMutex, s.Procs, s.BashPID, s.Host)
	logger.Add(logger.Disconnection, s.Host, s.Host, s.ID)
}

func newSession(s ssh.Session) (*Session, error) {
	host, _, err := net.SplitHostPort(s.RemoteAddr().String())
	if err != nil {
		return nil, err
	}

	return &Session{
		Session:   s,
		Host:      host,
		Path:      "/root",
		ProcMutex: &sync.Mutex{},
	}, nil
}

func (s *Session) spawnBash() error {
	p := &proc.Process{
		PPID: 1,
		User: s.Session.User(),
		Cmd:  "-bash",
		Args: []string{},
	}
	err := p.New(s.ProcMutex, s.Procs, s.Host, 1)
	if err != nil {
		return err
	}
	s.BashPID = p.PID

	return nil
}

func (s *Session) writeBanner(banner []byte) error {
	_, err := s.Session.Write([]byte("\x1b[H\x1b[2J\x1b[3J"))
	if err != nil {
		return err
	}

	_, err = s.Session.Write(banner)
	return err
}

func InitSession(s ssh.Session) (*Session, error) {
	session, err := newSession(s)
	if err != nil {
		return nil, err
	}

	f, err := session.copyJSON(configPATH+dirJSON, sessionPATH+session.Host+dirJSON)
	if err != nil {
		return nil, err
	}
	session.Dirs, err = filesystem.Parse(f)
	if err != nil {
		return nil, err
	}
	session.Entry = session.Dirs["root"]

	f, err = session.copyJSON(configPATH+procPATH+procJSON, sessionPATH+session.Host+procJSON)
	if err != nil {
		return nil, err
	}
	session.Procs, err = proc.Parse(f)
	if err != nil {
		return nil, err
	}

	session.ID, err = session.fetchID()
	if err != nil {
		return nil, err
	}

	err = session.spawnBash()
	if err != nil {
		return nil, err
	}

	logger.Add(logger.Connection, session.Host, session.Host, session.ID)

	banner, err := configs.ReadBanner()
	if err != nil {
		return nil, err
	}
	return session, session.writeBanner(banner)
}
