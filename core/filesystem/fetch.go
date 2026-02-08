package filesystem

import (
	"errors"
	"strings"
)

func Fetch(dirs map[string]*Entry, path string) (*Entry, error) {
	var entry *Entry
	paths := strings.Split(strings.TrimPrefix(path, "/"), "/")
	for i, name := range paths {
		if i == 0 {
			e := dirs[name]
			if e == nil {
				return nil, errors.New("No such file or directory")
			}
			entry = e
			continue
		}

		newEntry := entry.Children[name]
		if newEntry == nil {
			return nil, errors.New("No such file or directory")
		}

		entry = newEntry
	}

	return entry, nil
}
