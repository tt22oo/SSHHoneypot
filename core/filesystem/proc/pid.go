package proc

import "sort"

func newPID(procs map[int]*Process) int {
	var pids []int
	for _, process := range procs {
		pids = append(pids, process.PID)
	}
	sort.Ints(pids)

	if len(pids) < 1 {
		return 1
	}
	return pids[len(pids)-1] + 1
}
