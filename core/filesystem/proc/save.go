package proc

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(procs map[int]*Process, host string) error {
	f, err := os.Create(fmt.Sprintf("sessions/%s/procs.json", host))
	if err != nil {
		return err
	}

	return json.NewEncoder(f).Encode(procs)
}
