package log

import (
	"encoding/csv"
	"fmt"
	"os"
)

func initLog(host, id string) error {
	path := fmt.Sprintf("sessions/%s/%s", host, id)
	_, err := os.Stat(path)
	if err == nil {
		return nil // 이미 존재함
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path+"/logs.csv", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	return w.Write([]string{
		"time",
		"type",
		"data",
	})
}
