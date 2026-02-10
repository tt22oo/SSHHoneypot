## SSH Honeypot
[한국어](README.md) | [English](README-EN.md)  

An SSH honeypot written in Go, aiming to preserve the state (files, directories, etc. created by an attacker) even after a session is terminated.

## 1. Reason for Creation

Many SSH honeypots reset the environment when a session is closed and a new connection is made, even if the system appears to operate normally.
Attackers can use this characteristic to easily detect whether they are dealing with a honeypot.

This project was created to make such detection more difficult and to record attacker behavior in a more realistic environment.
It also aims to directly understand how honeypots work.

## 2. How It Works

1. When an attacker connects via SSH, a session is created.
2. State is recorded based on the connecting IP.
3. Files and directories per IP are recorded in JSON.
4. The saved state is maintained even after the session ends.
5. All login attempt information is logged.

## 3. Project Structure

### core/auth

* Authentication processing
* `log`: logging of authentication attempts

### core/commands

* Command-related processing
* `file`: file-related commands
* `system`: system-related commands

### core/session

* Session management and shell processing
* `handler`: session handling
* `shell`: command parsing and processing
* `stream`: input/output processing
* `log`: session logging

## 4. Configuration Method

The initial configuration file is in `configs/configs.json`.

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

* `listen`: You can set the port on which to run the SSH honeypot.
* `banner`: You can set the banner of the SSH honeypot.
* `max_delay`: You can specify the maximum random delay during login or command execution. (ms)
* `auth`: You can set whether authentication is enabled for the SSH honeypot.
* `username`: When authentication is enabled, you can set the username. (`*` allows all usernames.)
* `password`: When authentication is enabled, you can set the password. (`*` allows all passwords.)
* `arch`: You can set the architecture of the SSH honeypot. This is the architecture displayed by the honeypot system.
* `host_name`: You can set the Host Name of the SSH honeypot. This is displayed by the honeypot system.

## 5. Future Plans

* Implement more natural shell parsing
* `proc` implementation: partially implemented as of 2026-02-10
* Improve realism of command behavior
* Support more Linux commands
* Develop a web-based dashboard
* Session and attack log visualization features
