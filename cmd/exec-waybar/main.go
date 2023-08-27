package main

import (
	"context"
	"linux/pkg/logger"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
)

const (
	sway = "sway"
)

func main() {
	log.Infof("Starting EXEC-WAYBAR service")

	if os.Getenv("XDG_CURRENT_DESKTOP") != sway {
		log.Fatalf("Can be used only on %s", sway)
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	counter := 0

	for {
		select {
		case <-ctx.Done():
			log.Infof("Received interrupt signal! Aboring...")
			os.Exit(1)
		default:
			err := exec.Command("waybar").Run()
			if err != nil {
				logger.Warn(err.Error())
			}
			counter++
			if counter >= 1000 {
				logger.Fatal("Exceeded number of attempts")
				os.Exit(1)
			}
		}
	}
}
