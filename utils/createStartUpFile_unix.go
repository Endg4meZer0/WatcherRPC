//go:build linux

package wrpcUtils

import (
	"fmt"
	"os"
)

func CreateStartUpFile() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.WriteFile("~/.config/autostart/watcherrpc.desktop", []byte(fmt.Sprintf("[Desktop Entry]\nType=Application\nName=WatcherRPC\nExec=%v/app\nComment=WatcherRPC for the custom Discord activity based on running processes\nX-GNOME-Autostart-enabled=true", wd)), 0777)
	if err != nil {
		return err
	}

	return nil
}
