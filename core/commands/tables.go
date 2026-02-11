package commands

import (
	"honeypot/core/commands/file"
	"honeypot/core/commands/system"
	"honeypot/core/session"
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
}
