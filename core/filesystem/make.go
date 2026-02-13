package filesystem

import "errors"

func Make(dirs map[string]*Entry, e *Entry, t, name, data, host string) error {
	switch t {
	case File:
		e.Children[name] = &Entry{
			Type: File,
			Meta: &MetaData{
				Size: len(data),
			},
			Data: &data,
		}

		return Save(dirs, host)
	case Directory:
		e.Children[name] = &Entry{
			Type:     Directory,
			Children: make(map[string]*Entry),
		}

		return Save(dirs, host)
	}

	return errors.New("unknown type")
}
