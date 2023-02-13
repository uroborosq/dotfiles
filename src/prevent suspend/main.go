package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/exp/slices"
)

func main() {
	exec.Command("sudo", "systemctl", "unmask", "suspend.target").Run()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	isSuspendDisabled := false
	for {
		select {
		case <-sigs:
			fmt.Println("Interrupt signal recevied, restoring defaults...")
			if isSuspendDisabled {
				exec.Command("sudo", "systemctl", "unmask", "suspend.target").Run()
			}
			os.Exit(0)
		default:
			cmd := exec.Command("playerctl", "status")
			title, _ := cmd.Output()
			if slices.Equal(title, []byte("Playing\n")) {
				if !isSuspendDisabled {
					exec.Command("sudo", "systemctl", "mask", "suspend.target").Run()
					isSuspendDisabled = true
				}
			} else if isSuspendDisabled {
				exec.Command("sudo", "systemctl", "unmask", "suspend.target").Run()
				isSuspendDisabled = false
			}
			time.Sleep(5 * time.Second)
		}
	}
}
