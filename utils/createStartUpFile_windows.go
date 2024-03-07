//go:build windows

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
	err = os.WriteFile(os.Getenv("APPDATA")+"\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\WatcherRPC.ps1", []byte(fmt.Sprintf("cd \"%v\"\ntimeout 150\nstart .\\app.exe", wd)), 0777)
	if err != nil {
		return err
	}

	return nil
}
