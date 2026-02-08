package filesystem

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(dirs map[string]*Entry, host string) error {
	f, err := os.Create(fmt.Sprintf("sessions/%s/dirs.json", host))
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(dirs)
}
