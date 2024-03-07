//go:build windows

package wrpcUtils

import (
	"fmt"
	"os/exec"

	wrpcTypes "watcherrpc/app/types"
)

func ProcessCheck(currentHighestIndex *int, processes []wrpcTypes.ProcessToWatch) {
	for i := range len(processes) {
		out, err := exec.Command("tasklist", "/fi", fmt.Sprintf(`IMAGENAME eq %v`, processes[i].ProcessName)).Output()
		if err != nil {
			fmt.Printf("Warning: couldn't use the command \"tasklist\" for the process with the name of %v", processes[i].ProcessName)
			continue
		}
		if string(out[:5]) == "INFO:" {
			continue
		} else {
			*currentHighestIndex = i
			break
		}
	}
}
