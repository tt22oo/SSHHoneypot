package main

import (
	"honeypot/core/server"
)

func main() {
	go server.StartSSH() // start ssh honeypot
	select {}
}
