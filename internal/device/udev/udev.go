package udev

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func InstallSudoUdev() error {
	cwd, _ := os.Getwd()

	err := os.Chmod("scripts/udev/install_udev.sh", 0755)
	if err != nil {
		fmt.Println("no access")
		return err
	}

	cmd := exec.Command(
		"pkexec",
		filepath.Join(cwd, "scripts/udev/install_udev.sh"),
	)

	return cmd.Run()
}
