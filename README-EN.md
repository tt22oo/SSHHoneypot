## SSH Honeypot

[한국어](README.md) | [English](README-EN.md)

An SSH honeypot written in Go, designed with the goal of preserving the state (files, directories, etc. created by an attacker) even after a session ends.

## 1. Purpose
Many SSH honeypots reset their environment when a session ends and a new connection is made, even if the system appears to function normally.
Attackers can leverage this behavior to easily detect that they are interacting with a honeypot.

This project was created to make such detection more difficult and to record attacker behavior in a more realistic environment.
It also aims to help the developer directly understand how honeypots work internally.

## 2. How It Works
1. A session is created when an attacker connects via SSH.
2. State is tracked based on the client’s IP address.
3. Files and directories for each IP are recorded in JSON format.
4. The saved state is preserved even after the session ends.
5. All login attempts are logged.

## 3. Project Structure
### core/auth
- Authentication handling
- `log`: authentication attempt logging

### core/commands
- Command-related handling
- `file`: file-related commands
- `system`: system-related commands

### core/filesystem
- File-related processing
- `proc`: proc-related handling

### core/session
- Session management and shell handling
- `handler`: session handling
- `shell`: command parsing and processing
- `stream`: input/output handling
- `log`: session logging

## 4. Configuration
The initial configuration file is located at `configs/configs.json`.
```json
{
    "configs": {
        "listen": ":22",
        "banner": "test1234",
        "max_delay": 100,
        "auth": {
            "auth": true,
            "username": "root",
            "password": "password"
        }
    },
    "system": {
        "arch": "x86_64",
        "host_name": "server"
    }
}
```
- `listen`: Sets the port on which the SSH honeypot will run.
- `banner`: Sets the SSH honeypot banner.
- `max_delay`: Specifies the maximum random delay (in ms) applied during login or command execution.
- `auth`: Enables or disables SSH authentication.
- `username`: Username when authentication is enabled (`*` allows any username).
- `password`: Password when authentication is enabled (`*` allows any password).
- `arch`: Sets the architecture displayed by the honeypot system.
- `host_name`: Sets the hostname displayed by the honeypot system.

## 5. Future Plans

- Implement more natural shell parsing
- Further `proc` implementation (partially implemented as of 2026-02-10)
- Improve realism of command behavior
- Support more Linux commands
- Develop a web-based dashboard
- Add visualization for sessions and attack logs
