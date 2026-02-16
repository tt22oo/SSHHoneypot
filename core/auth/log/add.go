package log

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var AuthMu sync.Mutex

func Add(ip, username, password string) error {
	AuthMu.Lock()
	defer AuthMu.Unlock()

	log.Printf(" \033[32m[LOGIN]\033[0m %s (%s:%s)\r\n", ip, username, password)

	path := fmt.Sprintf("logs/%s.csv", time.Now().Format("2006-01-02"))
	err := initLog(path)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	return w.Write([]string{
		time.Now().Format(time.RFC3339),
		"login_attempt",
		ip,
		username,
		password,
	})
}
