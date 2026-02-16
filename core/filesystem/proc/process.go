package proc

import (
	"errors"
	"sync"
	"time"
)

type Process struct {
	PID       int       `json:"pid"`
	PPID      int       `json:"ppid"`
	User      string    `json:"user"`
	Cmd       string    `json:"cmd"`
	Args      []string  `json:"args"`
	StartTime time.Time `json:"start_time"`
}

func (p *Process) New(mu *sync.Mutex, procs map[int]*Process, host string) error {
	mu.Lock()
	defer mu.Unlock()

	pid := newPID(procs)
	procs[pid] = p

	p.PID = pid
	p.StartTime = time.Now()

	//log.Printf(" \033[36m[PROC]\033[0m New Process %s (PID=%d PPID=%d USER=%s CMD=%s)\r\n", host, p.PID, p.PPID, p.User, p.Cmd)
	return Save(procs, host)
}

func Delete(mu *sync.Mutex, procs map[int]*Process, pid int, host string) error {
	mu.Lock()
	defer mu.Unlock()

	p := procs[pid]
	if p == nil {
		return errors.New("not found pid")
	}

	delete(procs, pid)
	//log.Printf(" \033[36m[PROC]\033[0m Deleted Process %s (PID=%d)\r\n", host, p.PID)

	return Save(procs, host)
}
