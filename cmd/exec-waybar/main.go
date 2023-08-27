package main

import (
	"context"
	"os"
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

	for {
		select {
			case <-ctx.Done():
				log.Infof("Received interrupt signal! Aboring...")
				os.Exit(1)
			default:
				
		}
	}
}