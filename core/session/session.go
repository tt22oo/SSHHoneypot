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
	configPATH string = "configs"
	procPATH   string = "/proc"
	procJSON   string = "/procs.json"
	dirJSON    string = "/dirs.json"
)

func newDirs(host string) (*os.File, error) {
	path := fmt.Sprintf("sessions/%s", host)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}

	path += dirJSON
	_, err = os.Stat(path)
	if err != nil {
		data, err := os.ReadFile(configPATH + dirJSON)
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

func newProcs(host string) (*os.File, error) {
	path := fmt.Sprintf("sessions/%s", host)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return nil, err
	}

	path += procJSON
	_, err = os.Stat(path)
	if err != nil {
		data, err := os.ReadFile(configPATH + procPATH + procJSON)
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

func newSession(s ssh.Session) (*Session, *os.File, *os.File, error) {
	host, _, err := net.SplitHostPort(s.RemoteAddr().String())
	if err != nil {
		return nil, nil, nil, err
	}

	fdirs, err := newDirs(host)
	if err != nil {
		return nil, nil, nil, err
	}

	fprocs, err := newProcs(host)
	if err != nil {
		return nil, nil, nil, err
	}

	id, err := fetchID(host)
	if err != nil {
		return nil, nil, nil, err
	}

	return &Session{
		ID:        id,
		Session:   s,
		Host:      host,
		Path:      "/root",
		ProcMutex: &sync.Mutex{},
	}, fdirs, fprocs, nil
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

func InitSession(s ssh.Session) (*Session, error) {
	session, fdirs, fprocs, err := newSession(s)
	if err != nil {
		return nil, err
	}

	session.Dirs, err = filesystem.Parse(fdirs)
	if err != nil {
		return nil, err
	}
	session.Entry = session.Dirs["root"]

	session.Procs, err = proc.Parse(fprocs)
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
