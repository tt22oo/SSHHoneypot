package proc

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fetchProc(procs map[int]*Process) []string {
	var result []string
	for pid := range procs {
		result = append(result, fmt.Sprintf("%d", pid))
	}

	result = append(result, "cpuinfo")
	result = append(result, "meminfo")

	return result
}

func fetchCmdline(procs map[int]*Process, pid int) string {
	var result string
	result += procs[pid].Cmd
	for _, arg := range procs[pid].Args {
		result += " " + arg
	}

	return result
}

func fetchINFO(name string) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("configs/proc/%s.txt", name))
	if err != nil {
		return "", err
	}

	return string(data), err
}

func Fetch(procs map[int]*Process, path string) (string, error) {
	var result string
	paths := strings.Split(strings.Trim(path, "/"), "/")
	switch len(paths) {
	case 1:
		if path == "/proc" {
			plist := fetchProc(procs)
			for _, pid := range plist {
				result += pid + " "
			}

			result += "\r\n"
		}
	case 2:
		if paths[1] == "cpuinfo" || paths[1] == "meminfo" {
			return fetchINFO(paths[1])
		}
	case 3:
		if paths[0] == "proc" {
			if paths[2] == "cmdline" {
				pid, err := strconv.Atoi(paths[1])
				if err != nil {
					return "", err
				}

				result += fetchCmdline(procs, pid)
			}
		}
	default:
		return "", errors.New("not found")
	}

	return result, nil
}
