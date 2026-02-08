package shell

import (
	"fmt"
	"honeypot/core/configs"
	"honeypot/core/session"
	"honeypot/core/session/stream"
)

// print prompt
func writePrompt(s *session.Session) error {
	prompt := fmt.Sprintf("%s@%s:%s# ", s.Session.User(), configs.Cfg.System.HostName, s.Path)
	return stream.Output(s, prompt)
}
