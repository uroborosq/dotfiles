package main

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"

	"linux/pkg/logger"
)

const (
	sway  = "sway"
	limit = 1000
)

func main() {
	log.Infof("Starting EXEC-WAYBAR service")

	if os.Getenv("XDG_CURRENT_DESKTOP") != sway {
		log.Fatalf("Can be used only on %s", sway)
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	for range limit {
		select {
		case <-ctx.Done():
			log.Infof("Received interrupt signal! Aboring...")
			os.Exit(1)
		default:
			err := exec.Command("waybar").Run()
			if err != nil {
				logger.Warnf(err.Error())
			}
		}
	}
	logger.Fatalf("Exceeded number of attempts: %d", limit)
	os.Exit(1)
}
