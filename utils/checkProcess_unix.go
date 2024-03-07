//go:build linux

package wrpcUtils

import (
	"fmt"
	"os/exec"
	"strings"

	wrpcTypes "watcherrpc/app/types"
)

func ProcessCheck(currentHighestIndex *int, processes []wrpcTypes.ProcessToWatch) {
	for i := range len(processes) {
		out, err := exec.Command("ps", "-C", processes[i].ProcessName).Output()
		if err != nil {
			fmt.Printf("Warning: couldn't use the command \"ps\" for the process with the name of %v", processes[i].ProcessName)
			continue
		}
		if len(strings.Split(string(out), "\n")) == 1 {
			continue
		} else {
			*currentHighestIndex = i
			break
		}
	}
}
