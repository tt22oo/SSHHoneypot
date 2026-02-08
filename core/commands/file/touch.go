package file

import (
	"honeypot/core/filesystem"
	"honeypot/core/session"
)

func Touch(s *session.Session, args []string) (string, int) {
	data := ""
	s.Entry.Children[args[1]] = &filesystem.Entry{
		Type: filesystem.TypeFile,
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
