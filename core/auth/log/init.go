package log

import (
	"encoding/csv"
	"os"
)

func initLog(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	return w.Write([]string{
		"time",
		"type",
		"ip",
		"username",
		"password",
	})
}
