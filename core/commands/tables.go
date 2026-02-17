package commands

import (
	"honeypot/core/commands/file"
	"honeypot/core/commands/system"
	"honeypot/core/session"
	"log"
)

type command func(s *session.Session, args []string, pid int) (string, int)

var commandTables = make(map[string]command)

func add(name string, cmd command) {
	commandTables[name] = cmd
}

func Init() {
	add("cd", file.Cd)
	add("ls", file.Ls)
	add("cat", file.Cat)
	add("touch", file.Touch)
	add("mkdir", file.Mkdir)
	add("pwd", file.Pwd)

	add("id", system.Id)
	add("whoami", system.Whoami)
	add("uname", system.Uname)
	add("arch", system.Arch)
	add("clear", system.Clear)
	add("sleep", system.Sleep)
	add("ps", system.Ps)

	log.Printf(" \033[32m[SUCCESS]\033[0m Commands Loaded (%d)\r\n", len(commandTables))
}
