package proc

import (
	"encoding/json"
	"os"
)

func Parse(f *os.File) (map[int]*Process, error) {
	defer f.Close()

	var procs map[int]*Process
	err := json.NewDecoder(f).Decode(&procs)
	return procs, err
}
