## SSH Honeypot

[한국어](README.md) | [English](README-EN.md)

An SSH honeypot written in Go, aimed at maintaining the state such as files and directories created by an attacker even after the session ends.

## 1. Reason for making

Many SSH honeypots reset the environment when a session is closed and reconnected, even if the system is operating normally.
Attackers can easily detect honeypots by using these characteristics.

This project was created to make such detection more difficult and to record attacker behavior in a more realistic environment.
It also aims to directly understand how honeypots work.

## 2. How it works

1. When an attacker connects via SSH, a session is created.
2. State is recorded based on the connecting IP.
3. Files and directories per IP are recorded in JSON.
4. The saved state is maintained even after the session ends.
5. All login attempt information is recorded in logs.

## 3. Project structure

### core/server
- start server and management

### core/auth

- Authentication handling
- `log`: authentication attempt logging

### core/commands

- Command-related handling
- `file`: file-related commands
- `system`: system-related commands

### core/filesystem

- File-related handling
- `proc`: proc-related handling

### core/session

- Session management and shell handling
- `handler`: session handling
- `shell`: command parsing and handling
- `stream`: input/output handling
- `log`: session logging

## 4. Configuration method

**The default configuration file is in `configs/configs.json`.**

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

- `listen`: You can set the port to run the SSH honeypot.
- `banner`: You can set the banner of the SSH honeypot.
- `max_delay`: You can specify the maximum random delay during login or command execution. (ms)
- `auth`: You can set whether authentication is used for the SSH honeypot.
- `username`: If authentication is enabled, you can set the username. (Entering `*` allows all usernames.)
- `password`: If authentication is enabled, you can set the password. (Entering `*` allows all passwords.)
- `arch`: You can set the architecture of the SSH honeypot. This is the architecture displayed on the honeypot system.
- `host_name`: You can set the Host Name of the SSH honeypot. This is displayed on the honeypot system.

**`proc` related settings are in `configs/proc`.**

- `cpuinfo.txt`: You can set the information displayed in `/proc/cpuinfo` in the SSH honeypot.
- `meminfo.txt`: You can set the information displayed in `/proc/meminfo` in the SSH honeypot.
- `procs.json`:

```json
{
    "1": {
        "pid": 1,
        "ppid": 0,
        "user": "root",
        "cmd": "sleep",
        "args": ["100"],
        "start_time": "2026-02-09T03:43:58+09:00"
    }
}
```

- `pid`: You can set the pid of the process.
- `ppid`: You can set the pid of the parent process.
- `user`: You can set the user who owns the process.
- `cmd`: You can set the executed command.
- `args`: You can set the arguments passed when executing the process.
- `start_time`: You can set the time when the process started.

## 5. Future plans

- More natural shell parsing implementation
- ~~`proc` implementation~~
- Improving realism of command behavior
- Supporting more Linux commands
- Developing a web-based dashboard
- Session and attack log visualization features
