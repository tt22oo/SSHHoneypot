package main

import (
	"honeypot/core/server"
	"log"
)

func main() {
	err := server.StartSSH()
	if err != nil {
		log.Printf(" \033[31m[ERROR]\033[0m %s\r\n", err.Error())
		return
	}
}
