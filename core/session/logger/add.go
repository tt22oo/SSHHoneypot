package logger

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type DataType string

const (
	Connection DataType = "connection"
	Command    DataType = "command"
)

func Add(t DataType, data, host, id string) error {
	err := initLog(host, id)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fmt.Sprintf("sessions/%s/%s/logs.csv", host, id), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	log.Printf(" \033[32m[SESSION]\033[0m New %s: (%s) (host=%s id=%s)\r\n", t, data, host, id)

	return w.Write([]string{
		time.Now().Format(time.RFC3339),
		string(t),
		data,
	})
}
