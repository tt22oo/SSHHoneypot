package file

import (
	"honeypot/core/filesystem"
	"honeypot/core/filesystem/proc"
	"honeypot/core/session"
)

func Touch(s *session.Session, args []string, pid int) (string, int) {
	defer proc.Delete(s.ProcMutex, s.Procs, pid, s.Host)

	data := ""
	s.Entry.Children[args[1]] = &filesystem.Entry{
		Type: filesystem.File,
		Meta: &filesystem.MetaData{
			Size: len(data),
		},
		Data: &data,
	}

	err := filesystem.Save(s.Dirs, s.Host)
	if err != nil {
		return "", 1
	}

	return "", 0
}
