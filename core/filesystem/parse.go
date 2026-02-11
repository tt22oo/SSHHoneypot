package filesystem

import (
	"encoding/json"
	"os"
)

type MetaData struct {
	Size int `json:"size"`
}

type Entry struct {
	Type     string            `json:"type"`
	Meta     *MetaData         `json:"metadata"`
	Children map[string]*Entry `json:"children"`
	Data     *string           `json:"data"`
}

var (
	File      = "file"
	Directory = "directory"
)

func Parse(f *os.File) (map[string]*Entry, error) {
	defer f.Close()

	var dirs map[string]*Entry
	err := json.NewDecoder(f).Decode(&dirs)
	return dirs, err
}
